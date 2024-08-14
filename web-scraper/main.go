package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main()  {
	//Instantiate collector
	c := colly.NewCollector(
		//This should  visit only the domains listed here 
		colly.AllowedDomains("go.dev/doc/effective_go", "go.dev/ref/spec"),
	)

	// extracts the URL from elements with href attribute
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		//Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page, only those specified in allowedDomains

		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://go.dev/doc/effective_go
	c.Visit("https://go.dev/doc/effective_go/")
}