package scrapper

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

type infoDict struct {
	id string
	title string
	corpname string
	location string
	content string
}

func Scrape(term string) {
	baseURL := "https://kr.indeed.com/jobs?q="+ term +"&limit=50"
	pages := getPages(baseURL)
	var totalDict []infoDict
	mainC := make(chan []infoDict)
	for i:=0;i<pages;i++ {
		go getParsedPage(i, baseURL, mainC)
	}
	for i:=0;i<pages;i++ {
		jobDict := <- mainC
		totalDict = append(totalDict, jobDict...)
	}
	writeJobs(totalDict)
}

func getPageLink(page int, url string) string {
	pageURL := url + "&start=" + strconv.Itoa(page*50)
	return pageURL
}

func getParsedPage(page int, base_url string, mainC chan<-[]infoDict){
	url := getPageLink(page, base_url)
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
	title := CleanText(s.Find(".title > a").Text())
	corpname := CleanText(s.Find(".sjcl > div:nth-child(1) > span").Text())
	location := CleanText(s.Find(".sjcl > span").Text())
	content := CleanText(s.Find(".summary").Text())
	c<-infoDict{
		id : "https://kr.indeed.com" + id,
		title : title,
		corpname: corpname,
		location: location,
		content: content,
	}
}

func getPages(url string) int {
	resp, err := http.Get(url)
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

func CleanText(str string) string {
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