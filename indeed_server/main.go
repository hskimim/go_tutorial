package main

import (
	"github.com/hskimim/go_tutorial/scrapper"
	"github.com/labstack/echo"
	"os"
	"strings"
)

const fileName string = "job.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	//scrapper.Scrape("python")
}

func handleHome(c echo.Context) error {
//	return c.String(http.StatusOK, "Hello, World!")
	return c.File("home.html")
}

func handleScrape(c echo.Context) error{
	//	return c.String(http.StatusOK, "Hello, World!")
	defer os.Remove(fileName)
	term := strings.ToLower(scrapper.CleanText(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(fileName, fileName)
//	return c.File("home.html")
}