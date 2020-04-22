package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	//	"reflect"
)

type infoDict struct {
	id string
	title string
	corpname string
	location string
	content string
}

func main() {
	start := time.Now()
	contentLs := extractContent("job.csv")
	wordCountMap := WordCount(contentLs)
	writeWordCount("content_wc.csv", wordCountMap)
	// Code to measure
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)
}

func readCSV(filename string, c chan []infoDict){
	var data []infoDict
	csvFile, err := os.Open(filename)
	checkErr(err)
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	checkErr(err)
	for _, line := range lines {
		tmp := infoDict{
			id:       line[0],
			title:    line[1],
			corpname: line[2],
			location: line[3],
			content:  line[4],
		}
		data = append(data, tmp)
	}
	c <-data
}

func extractContent(filename string) []string {
	c := make(chan []infoDict)
	go readCSV(filename, c)
	data := <- c
	contentContainer := []string{}
	for _, line := range data {
		contentContainer = append(contentContainer, line.content)
	}
	return contentContainer
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln("ERROR : ", err)
	}
}

func preprocessText(s string) string {
	lowString := strings.ToLower(s)
	reg, err := regexp.Compile("[^a-zA-Z0-9가-힣]+")
	checkErr(err)
	processedString := reg.ReplaceAllString(lowString, "")
	return processedString
}

func TotalWordCount(contentLs []string) map[string]int {
	c := make(chan map[string]int)
	wcSlice := []map[string]int{}
	for _, content := range contentLs {
		go WordCount(content, c)
		wc := <- c
		wcSlice = append(wcSlice,wc)
	}

}

func WordCount(s string, c chan map[string]int) {
	wordCountMap := make(map[string]int)
	words := strings.Fields(content) //split sentence by " "
	for _, word := range words {
		processedText := preprocessText(word)
		wordCountMap[processedText]++
	}
	c <- wordCountMap
}
func get_keys(mymap map[string]int) []string {
	keys := make([]string, len(mymap))
	i := 0

	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func writeWordCount(filename string, jobs map[string]int) {
	file, err := os.Create(filename)
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"word", "count"}
	wErr := w.Write(headers)
	checkErr(wErr)

	keyLs := get_keys(jobs)
	for _, word := range keyLs{
		wc := strconv.Itoa(jobs[word])
		jobSlice := []string{word, wc}
		jbErr := w.Write(jobSlice)
		checkErr(jbErr)
	}
}