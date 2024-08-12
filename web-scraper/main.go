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
	c.OnHTML("a[href]", func(e *colly.HTMLElement)  {
		link := e.Attr("href")

		//Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})
	
}