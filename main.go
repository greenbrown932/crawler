package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

type crawler struct {
	url        string
	urlStatus  string
	visited    bool
	rawHtml    string
	parsedHtml *html.Node
	DNS        string
	queue      []string
}

func fetchURL(c *crawler) {
	// Fetch the URL and update the crawler's status and raw HTML
	resp, err := http.Get(c.url)
	if err != nil {
		c.urlStatus = "error"
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.urlStatus = "error"
		return
	}
	c.rawHtml = string(body)
	c.visited = true
	c.DNS = resp.Request.URL.Hostname()
}

func parse(c *crawler) {
	// Parse the raw HTML and update the crawler's status and parsed HTML
	// Implement parsing logic here
	c.parsedHtml, err := html.Parse(c.rawHtml)
	if err != nil {
		c.urlStatus = "error"
		return
	}

}

func crawl(c *crawler) {
	// Fetch the URL and parse the HTML

}

func main() {
	fmt.Println("Hello, World!")

	c := crawler{
		url:       "https://example.com",
		urlStatus: "pending",
		visited:   false,
		rawHtml:   "",
		DNS:       "example.com",
		queue:     []string{},
	}
	fmt.Println(c)
}
