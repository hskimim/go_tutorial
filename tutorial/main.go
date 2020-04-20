package tutorial

import (
	"errors"
	"fmt"
	"github.com/hskimim/go_tutorial/tutorial/accounts"
	"github.com/hskimim/go_tutorial/tutorial/mydict"
	"net/http"
	"strings"
	"time"
)

var errResponse = errors.New("That web-site cannot connect")

// data type practice
func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, upper_name string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	upper_name = strings.ToUpper(name)
	return
}

func repeat(names ...string) {
	fmt.Println(names)
}

func addNumebers(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIdrink(age int) bool {
	if KoreanAge := age + 2; KoreanAge < 20 {
		return false
	}
	return true
}

func canIdrinkSwitch(age int) bool {
	KoreanAge := age + 2
	switch {
	case KoreanAge < 20:
		return false
	case KoreanAge >= 20:
		return true
	}
	return false
}

func pointerTest() (int, int) {
	a := 2
	b := &a
	a = 4
	return a, *b
	// (4,4)
}

func sliceExample() []string {
	slice := []string{"hello", "world"}
	new_slice := append(slice, "hyunsikkim")
	return new_slice
}

func mapExample() map[string]string {
	dict_ := map[string]string{"key": "value"}
	return dict_
}

type hyunsikkim struct {
	name  string
	age   int
	hobby []string
}

func structExample() string {
	hobby_ls := []string{"coding", "reading"}
	hskimim := hyunsikkim{name: "hyunsikkim", age: 23, hobby: hobby_ls}
	return hskimim.name
}

func main() {
	length, upper_name := lenAndUpper("hyunsikkim")
	fmt.Println(length, upper_name)

	repeat("hello", "my", "name", "is", "hyunsikkim")
	fmt.Println(addNumebers(12, 3, 4, 5, 6))

	fmt.Println(canIdrink(24))
	fmt.Println(canIdrinkSwitch(23))
	fmt.Println(pointerTest())
	fmt.Println(sliceExample())
	fmt.Println(mapExample()["key"])
	fmt.Println(structExample())
}

// ===================================================================================================================//

// bank account simulation (refer account)
func main() {
	account := accounts.NewAccount("hskimim")
	account.Deposit(10)
	err := account.Withdraw(100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}

// ===================================================================================================================//
// mapper manipulation practice (refer mydict)
func main() {
	dictionary := mydict.Dictionary{"first": "first word"}
	value, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}

// ===================================================================================================================//
// slow(vanila) urlchecker

func main() {
	dict_ := make(map[string]string)
	urls := []string{
		"https://google.com",
		"https://naver.com",
		"https://daum.net",
	}
	for _, url := range urls {
		err := hitURL(url)
		switch err {
		case nil:
			dict_[url] = "COMPLETE"
		case errResponse:
			dict_[url] = "ERROR"
		}
	}
	for k, v := range dict_ {
		fmt.Println(k, v)
	}
}

func hitURL(url string) error {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errResponse
	}
	return nil
}

// ===================================================================================================================//
// goroutine practice

func IsSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}

func main() {
	c := make(chan string)
	names := []string{"hskimim", "hrlee"}
	for _, name := range names {
		go IsSexy(name, c)
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(<-c)
	}
}

// ===================================================================================================================//
// fast(with goroutine) urlchecker

type result struct {
	url    string
	status string
}

func main() {
	c := make(chan result)
	status_dict := make(map[string]string)
	urls := []string{
		"https://google.com",
		"https://naver.com",
		"https://daum.net",
		"https://instagram.com",
		"https://github.com",
		"https://academy.nomadcoders.co/",
		"https://facebook.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		status_dict[result.url] = result.status
	}

	for key, value := range status_dict {
		fmt.Println(key, value)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get(url)
	status := "COMPLETE"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAIL"
		c <- result{url: url, status: status}
	} else {
		c <- result{url: url, status: status}
	}
}
