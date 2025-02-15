package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

const baseURL = "https://gobyexample.com/"

func main() {
	urls := make(map[string]string)
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response status:", r.StatusCode)
	})
	c.OnHTML("ul li a", func(e *colly.HTMLElement) {
		title := e.Text
		link := baseURL + e.Attr("href")
		fmt.Printf("%s -> %s\n", title, link)
		urls[title] = link
	})

	for title, link := range urls {
		fmt.Println("fatch: " + title + " -> " + link)
	}

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request failed:", err)
	})

	err := c.Visit(baseURL)
	if err != nil {
		log.Fatal(err)
	}
}
