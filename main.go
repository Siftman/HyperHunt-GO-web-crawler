package main

import (
	"fmt"

	"github.com/siftman/HyperHunt-GO-web-crawler/pkg/crawler"
	"github.com/siftman/HyperHunt-GO-web-crawler/pkg/fileops"
	"github.com/siftman/HyperHunt-GO-web-crawler/pkg/utils"
)

func main() {
	baseURL := "https://khanesaat.com/"
	status, url := crawler.SitemapStatusCheck(baseURL)

	if status == 200 {
		urls := crawler.SitemapScanner(url)
		fileops.WriteCSV("raw_links.csv", urls)
		frequent_path := utils.Frequency_finder(urls)
		proper_urls := utils.Filter_csv(urls, frequent_path)
		fileops.WriteCSV("proper_urls.csv", proper_urls)

		for _, url := range proper_urls {
			fmt.Println(crawler.Crawler(url))
		}
	}
}
