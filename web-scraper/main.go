package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main()  {
	
	//Instantiate collector
	c := colly.NewCollector(

		//This should  visit only the domains listed here 
		colly.AllowedDomains("go.dev/doc/effective_go", "go.dev/ref/spec "),
	)


	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
	
}