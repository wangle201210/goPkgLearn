package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func ExampleScrape() {
	client := &http.Client{}
	response, err := client.Get("https://wenku.baidu.com/view/fb15d08fec630b1c59eef8c75fbfc77da36997f0.html")
	if err != nil {
		log.Printf("client.Do err (%+v)", err)
		return
	}
	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
	fmt.Printf("%+v",response.Body)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#reader-container").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		fmt.Printf("Review %d: %s - %s\n", i, band)
	})
}

func main() {
	ExampleScrape()
}