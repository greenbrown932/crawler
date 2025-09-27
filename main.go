package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type crawler struct {
	url        string
	urlStatus  string
	visited    bool
	rawHtml    string
	parsedHtml *html.Node
	DNS        string
}

func fetchURL(c *crawler) {
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

func parse(c *crawler) error {
	var err error
	c.parsedHtml, err = html.Parse(strings.NewReader(c.rawHtml))
	if err != nil {
		c.urlStatus = "error"
		return err
	}
	return nil
}

func crawl(c *crawler) {
	// Fetch the URL and parse the HTML
	fetchURL(c)
	parse(c)
	fmt.Println("URL:", c.url)
	fmt.Println("DNS:", c.DNS)
	fmt.Println("Visited:", c.visited)
	fmt.Println("Status:", c.urlStatus)
	fmt.Println("Raw HTML:", c.rawHtml)
	fmt.Println("Parsed HTML:", c.parsedHtml)

}

func main() {

	urls := []string{"https://google.com", "https://example.com"}

	for _, url := range urls {
		c := crawler{
			url:        url,
			urlStatus:  "pending",
			visited:    false,
			rawHtml:    "",
			DNS:        "example.com",
			parsedHtml: &html.Node{},
		}

		crawl(&c)
	}
}
