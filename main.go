package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"
type infoDict struct {
	id string
	title string
	corpname string
	location string
	content string
}

func main() {
	pages := getPages()
	var totalDict []infoDict
	mainC := make(chan []infoDict)
	for i:=0;i<pages;i++ {
		go getParsedPage(i, mainC)
	}
	for i:=0;i<pages;i++ {
		jobDict := <- mainC
		totalDict = append(totalDict, jobDict...)
	}
	writeJobs(totalDict)
}

func getPageLink(page int) string {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	return pageURL
}

func getParsedPage(page int, mainC chan<-[]infoDict){
	url := getPageLink(page)
	fmt.Println(url, " is crawled now...")
	resp, err := http.Get(url)
	var jobDict []infoDict
	c := make(chan infoDict)
	checkRes(resp)
	checkErr(err)
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	div := doc.Find(".jobsearch-SerpJobCard")
	div.Each(func(i int, s *goquery.Selection) {
		go extractJob(s, c)
		job := <- c
		jobDict = append(jobDict, job)
		})
	mainC <- jobDict
}

func extractJob(s *goquery.Selection, c chan<-infoDict) {
	id, _ := s.Find("h2 > a").Attr("href")
	title := cleanText(s.Find(".title > a").Text())
	corpname := cleanText(s.Find(".sjcl > div:nth-child(1) > span").Text())
	location := cleanText(s.Find(".sjcl > span").Text())
	content := cleanText(s.Find(".summary").Text())
	c<-infoDict{
		id : "https://kr.indeed.com" + id,
		title : title,
		corpname: corpname,
		location: location,
		content: content,
	}
}

func getPages() int {
	resp, err := http.Get(baseURL)
	checkErr(err)
	checkRes(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	return getTotalPages(doc)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln("There was error in http.Get")
	}
}

func checkRes(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln("The status code is not valid ", resp.StatusCode)
	}
}

func getTotalPages(doc *goquery.Document) int {
	totalPages := 0
	doc.Find("#resultsCol > div.pagination").Each(func(i int, s *goquery.Selection) {
		totalPages = s.Find("a").Length()
	})
	return totalPages
}

func cleanText(str string) string {
	clean_str := strings.Join(strings.Fields(strings.TrimSpace(str))," ")
	clean_str = strings.Replace(clean_str, "  ", " ", -1)
	return clean_str
}

func writeJobs(jobs []infoDict) {
	file, err := os.Create("job.csv")
	checkErr(err)
	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"url", "title", "corpname", "location", "content"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs{
		jobSlice := []string{job.title, job.corpname, job.location, job.content}
		jbErr := w.Write(jobSlice)
		checkErr(jbErr)
	}
}