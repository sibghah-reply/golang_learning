package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

//todo: complete the scrapePage function and then the scrapeallpages function
//next: write a function that asks gpt to select each of the selectors, - gpt should not do the scraping, that will consume too many resources and take too long
//from thsi function either gpt will be asked for selectors or these selectors will be retrieved from the user
//later add functions that check the validity of the url etc

func retriveHTMLContent(selectors map[string][]string, url string, fileName string, waitForSelector bool) (playwright.Page, error) {
	mainContentSelectors := selectors["mainContentSelectors"]
	headingSelectors := selectors["headingSelectors"]
	headingContentSelectors := selectors["headingContentSelectors"]

	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("Could not start Playwright: %v", err)
	}

	browser, err := pw.Chromium.Launch(
		playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(true),
		},
	)

	if err != nil {
		log.Fatalf("Could not launch the browser: %v", err)
	}

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Could not open a new page: %v", err)
	}

	if _, err = page.Goto(url); err != nil {
		log.Fatalf("Could not visit the desired page: %v", err)
	}

	fmt.Println(mainContentSelectors, headingSelectors, headingContentSelectors)
	return page, err
}

func getMainContent(mainSelectors []string, page playwright.Page) {
	mainTextContent := []string{}

	for _, selector := range mainSelectors {

		mainContentMatchesLocator := page.Locator(selector)
		textContents, err := mainContentMatchesLocator.AllInnerTexts()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		mainTextContent = append(mainTextContent, textContents...)

	}

	fmt.Println(mainTextContent)

}

func main() {
	// Define a map of selectors
	selectors := map[string][]string{
		"mainContentSelectors":    {"#main", ".content"},       // Example main content selectors
		"headingSelectors":        {".heading", ".title"},      // Example heading selectors
		"headingContentSelectors": {".subheading", ".summary"}, // Example heading content selectors
	}

	// Example URL and filename
	url := "http://amazon.com"
	fileName := "exampleFile"

	// Boolean for waitForSelector
	waitForSelector := true

	// Call the scrapePage function with the necessary arguments - should this go somewhere else?
	page, err := retriveHTMLContent(selectors, url, fileName, waitForSelector)
	if err != nil {
		log.Fatalf("Could not retrieve page content: %v", err)
	}

	getMainContent([]string{"body"}, page)

}
