package main

import (
	"fmt"
	"log"
	"os"

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
		colly.AllowedDomains("www.albertsons.com"),
	)
	// var items []item

	// c.OnHTML("product-item-al", func(e *colly.HTMLElement) {
	// 	item := item{
	// 		price: e.ChildText(".product-price"),
	// 	}
	// 	items = append(items, item)
	// })
	file, err := os.Create("output.html")
	if err != nil{
		fmt.Println("error creating file", err)
		return	
	}
	defer file.Close()


c.OnHTML("body",func(e *colly.HTMLElement) {
	_, err = file.WriteString(e.Text)
	if err != nil{
fmt.Println("failed to write to file:", err)
return	
	}
fmt.Println("Wrote to file", file.Name())

})

c.OnRequest(func(r *colly.Request) {
	fmt.Println("visiting",r.URL)
})

c.OnResponse(func(r *colly.Response) {
	fmt.Println("response received", r.Request.URL )
})

c.OnError(func(r *colly.Response, err error) {
	if err != nil {
	log.Println("whoops!",err)
	}
})

c.Visit("https://www.albertsons.com/shop/search-results.html?q=egg")

fmt.Println("scrapping done /n", )
}
