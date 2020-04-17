package main

import (
	"fmt"
	"strings"
)

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
