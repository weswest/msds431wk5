package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"encoding/json"

	"net/url"

	"github.com/gocolly/colly"
)

type URL struct {
	URL string `json:"url"`
}

type Website struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Function to read URLs from a .jl file

func readURLs(filename string) ([]string, error) {
	// Open the .jl file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new Scanner for the file
	scanner := bufio.NewScanner(file)

	// Create an empty slice to hold the URLs
	var urls []string

	// Loop over all lines in the file
	for scanner.Scan() {
		// Get the next line
		line := scanner.Text()

		// Unmarshal the JSON line into a URL object
		var u URL
		if err := json.Unmarshal([]byte(line), &u); err != nil {
			return nil, err
		}

		// Append the URL to the urls slice
		urls = append(urls, u.URL)
	}

	// Check for errors from scanner
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

// Pulls the URL and returns the page's raw HTML, title, and text.
func pullURL(url string) ([]byte, string, string, error) {
	var body []byte
	var title string
	var text string

	c := colly.NewCollector()

	// Setup rate limiting
	c.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "*",
		// Set a delay between requests
		Delay: 1 * time.Second,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	// Actually get the HTML context
	c.OnResponse(func(r *colly.Response) {
		body = r.Body
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		text = e.Text
	})

	err := c.Visit(url)
	if err != nil {
		return nil, "", "", err
	}

	return body, title, text, nil
}

// Writes the raw HTML to a file.
func writeToHTML(givenUrl string, htmlData []byte) error {
	parsedURL, err := url.Parse(givenUrl)
	if err != nil || parsedURL.Host == "" || parsedURL.Scheme == "" {
		return fmt.Errorf("Invalid URL: %v", givenUrl)
	}

	// Get the last segment of the URL path.
	urlPath := path.Base(givenUrl)

	// Create the filename.
	fileName := fmt.Sprintf("wikipages/%s.html", urlPath)

	// Write the HTML data to the file.
	err = os.WriteFile(fileName, htmlData, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Appends the URL, title, and text to the .jl file.
func writeToJL(givenUrl string, title string, text string) error {
	parsedURL, err := url.Parse(givenUrl)
	if err != nil || parsedURL.Host == "" || parsedURL.Scheme == "" {
		return fmt.Errorf("Invalid URL: %v", givenUrl)
	}

	website := Website{
		URL:   givenUrl,
		Title: title,
		Text:  text,
	}

	f, err := os.OpenFile("goItems.jl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	itemJSON, err := json.Marshal(website)
	if err != nil {
		return err
	}

	_, err = f.WriteString(string(itemJSON) + "\n")
	if err != nil {
		return err
	}

	return nil
}

// If the startFresh flag is true, attempts to delete the "wikipages" directory and the "goItems.jl" file.
func startFresh(startFresh bool) error {
	if startFresh {
		err := os.RemoveAll("wikipages")
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		err = os.Mkdir("wikipages", 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}

		err = os.Remove("goItems.jl")
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}

func main() {
	err := startFresh(true)
	if err != nil {
		log.Fatalf("Error in startFresh: %v", err)
	}

	// Read URLs from the .jl file
	urls, err := readURLs("testURLs.jl")
	if err != nil {
		log.Fatalf("Failed to read URLs: %s", err)
	}

	for i, nextUrl := range urls {
		fmt.Printf("Now attempting to download item %d of %d: %s\n", i+1, len(urls), nextUrl)
		body, title, text, err := pullURL(nextUrl)
		if err != nil {
			log.Fatalf("Error pulling URL %s: %v", nextUrl, err)
		}

		err = writeToHTML(nextUrl, body)
		if err != nil {
			log.Fatalf("Error writing to HTML: %v", err)
		}

		err = writeToJL(nextUrl, title, text)
		if err != nil {
			log.Fatalf("Error writing to JL: %v", err)
		}
	}
}
