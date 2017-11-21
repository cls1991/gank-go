package core

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func Extract(url string) ([]string, error) {
	targetUrl := url
	var links []string
	for {
		fmt.Printf("Fetch url: %s\n", targetUrl)
		doc, err := goquery.NewDocument(targetUrl)
		if err != nil {
			return nil, err
		}
		// detect pic link
		val, exist := doc.Find("div.outlink").Find("img").Attr("src")
		if !exist {
			return nil, err
		}
		if strings.HasSuffix(val, ".jpg") || strings.HasSuffix(val, ".png") || strings.HasSuffix(val, "jpeg") {
			links = append(links, val)
		}
		// jump to pre page
		pre, exist := doc.Find("[style=\"text-align: right\"]").Find("A").Attr("href")
		if !exist {
			break
		}
		targetUrl = url + pre
	}
	fmt.Println("Done.")

	return links, nil
}
