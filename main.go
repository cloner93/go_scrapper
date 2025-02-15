package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gocolly/colly"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const baseURL = "https://gobyexample.com/"
const saveDir = "example_dir/"

func main() {
	// Create a new Colly collector
	c := colly.NewCollector()

	// Ensure the output directory exists
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create directory:", err)
	}

	// Scrape main page and visit each example link
	c.OnHTML("ul li a", func(e *colly.HTMLElement) {
		// title := e.Text
		link := baseURL + e.Attr("href")

		// Visit each example page
		e.Request.Visit(link)
	})
	// hold docs and code in map
	dataArr := []string{}

	// Scrape example pages and extract docs + multi-line code
	c.OnHTML("table", func(e *colly.HTMLElement) {

		// Extract Docs (First <td class="docs"> with content)
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			doc := strings.TrimSpace(row.ChildText("td.docs"))
			code := strings.TrimSpace(row.ChildText("td.code pre"))

			// Save the first non-empty docs
			if doc != "" || code != "" {
				if doc == "" {
					dataArr = append(dataArr, "\n")
				} else {
					doc = "// " + strings.ReplaceAll(doc, "\n", " ")
					dataArr = append(dataArr, "\n"+doc+"\n")
				}

				dataArr = append(dataArr, code+"\n")
			}
		})

		// Get page title from URL for filename
		pageTitle := e.Request.URL.Path // e.g., "/hello-world"
		fileName := formatFileName(strings.Trim(pageTitle, "/")) + ".go"

		// Save the file
		saveToFile(fileName, dataArr)

		// Clear data
		dataArr = dataArr[:0]
	})

	// Debug: Print errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request failed:", err)
	})

	// Start scraping the main page
	err = c.Visit(baseURL)
	if err != nil {
		log.Fatal(err)
	}
}

// Function to format filename properly
func formatFileName(title string) string {
	re := regexp.MustCompile(`[-_]`)
	title = re.ReplaceAllString(title, " ")
	title = cases.Title(language.English).String(title) // Proper capitalization
	return strings.ReplaceAll(title, " ", "-")
}

// Function to save docs + code to a .go file
func saveToFile(fileName string, data []string) error {
	// File path
	filePath := filepath.Join(saveDir, fileName)

	// check if file exists return
	if _, err := os.Stat(filePath); err == nil {
		return nil
	}

	// Join data into a single string
	content := strings.Join(data, "")

	// Write to file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Println("Error writing file:", err)
		return err
	}

	fmt.Println("Saved:", fileName)
	return nil
}
