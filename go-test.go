package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// this is the typing for each scraped item to follow
type item struct {
	name        string
	price       int
	img         string
	description string
	size        string
}

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains("www.albertsons.com"),
	)
	var items []item
	c.OnHTML("product-item-al", func(e *colly.HTMLElement) {
		item := item{
			price: e.ChildText(".product-price"),
		}
		items = append(items, item)
	})

	fmt.Println()
}
