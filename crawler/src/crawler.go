package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

//next: push code so far to github
//maybe add a readme to github
//read the article about the distributed crawler and potentially add more features to it

type Crawler struct {
	seedUrls    []string
	urlQueue    []string
	visitedUrls map[string]bool //maps alone are not safe for concurrent use, so this cannot
	mutex       sync.Mutex      // Protects shared resources
	wg          sync.WaitGroup  // Wait for all goroutines to finish
	urlChan     chan string
	stopped     bool
}

// equivalent to a class constructor
func NewCrawler(seedUrls []string) *Crawler {
	return &Crawler{
		seedUrls:    seedUrls,
		urlQueue:    make([]string, len(seedUrls)),
		visitedUrls: make(map[string]bool),
		urlChan:     make(chan string, 100),
		stopped:     false,
	}
}

// Download the content of the page
// inside the first parens is a method receiver: specifies the type that the function can be called on, similar to having a class and functions which can then only be called on an instance of that class, i.e. self
// second set of parens is parameters, third set has required return types, both are required, nil will be returned in the case of no error
func (c *Crawler) downloadURL(pageURL string) (string, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer func() {
		resp.Body.Close()
	}()
	// runs when the function exits, ensures that the response body is closed, released resources, whether or not there is an error
	//deferring is usually for external resources

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch page: %s", pageURL)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	html, err := doc.Html() //HTML is a method that belongs to the goquery.Document type
	if err != nil {
		return "", err
	}

	return html, nil
}

// Extract linked URLs from the page
func (c *Crawler) getLinkedURLs(pageURL string, html string) []string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}

	var urls []string
	baseURL, err := url.Parse(pageURL)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) { //more efficient to use a pointer in this case, it's common practice to use pointers for types of data that represent collections, data structures or anything that requires modifications or has as significant amount of data
		href, exists := s.Attr("href")
		if exists {
			parsedURL, err := url.Parse(href)
			if err == nil {
				absURL := baseURL.ResolveReference(parsedURL).String()
				urls = append(urls, absURL)
			}
		}
	})
	return urls
}

// Check if URL has already been visited or queued
func (c *Crawler) addURLToVisit(u string) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.visitedUrls[u] {
		c.visitedUrls[u] = true
		return true
	}
	return false
}

// Crawl the provided URL concurrently
// : This is a channel of type string, used for sending newly discovered URLs back to the main goroutine or another part of the program.
func (c *Crawler) crawl(pageURL string) {
	c.wg.Add(1)
	fmt.Println("added")
	defer func() {
		c.wg.Done()
		fmt.Println("done")
	}()

	html, err := c.downloadURL(pageURL)
	if err != nil {
		log.Println(err)
		return
	}

	linkedURLs := c.getLinkedURLs(pageURL, html)
	for _, link := range linkedURLs {
		if c.addURLToVisit(link) {
			c.urlChan <- link // Send newly discovered URLs to the channel
		}
	}
}

func (c *Crawler) crawlAllURLs() {
	// Seed the queue with initial URLs
	for _, seedURL := range c.seedUrls {
		c.urlQueue = append(c.urlQueue, seedURL)

		go c.crawl(seedURL)
	}

	go func() {
		defer func() {

		}()
	}()

	// Start a separate goroutine to handle crawling of discovered URLs
	for url := range c.urlChan {
		fmt.Println("Discovered:", url)

		go c.crawl(url) // Crawl discovered URLs concurrently
		// Wait for all crawlers to finish before closing the channel
		go func() {
			c.wg.Wait() // Wait for all crawling tasks to finish

			//dont close a closed channel
			c.mutex.Lock()
			if !c.stopped {
				close(c.urlChan)
				c.stopped = true
			}
			c.mutex.Unlock()
		}()
	}

}

func main() {
	crawler := NewCrawler([]string{"https://stackoverflow.com/questions/70908948/channel-hangs-probably-not-closing-at-the-right-place"})
	crawler.crawlAllURLs()
}
