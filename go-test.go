package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

// this is the typing for each scraped item to follow
type item struct {
	name        string
	price       string
	img         string
	description string
	size        string
}

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains("www.albertsons.com"),
	)
	// var items []item

	// c.OnHTML("product-item-al", func(e *colly.HTMLElement) {
	// 	item := item{
	// 		price: e.ChildText(".product-price"),
	// 	}
	// 	items = append(items, item)
	// })



c.OnRequest(func(r *colly.Request) {
	fmt.Println("visiting",r.URL)
})

c.OnResponse(func(r *colly.Response) {
	fmt.Println("response received", r.Request.URL )
})

c.OnError(func(r *colly.Response, err error) {
	log.Println("whoops!",err)
})

c.Visit("https://www.albertsons.com/shop/search-results.html?q=egg")

c.Wait()
fmt.Println("scrapping done /n", document)
}
