package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	fName := "WikiLinks.txt"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Could not create file: ", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	// Find and print all links
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {

		links := e.ChildAttrs("a", "href")
		for _, link := range links {
			file.WriteString(link + "\n")
		}

	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
	defer writer.Flush()
}
