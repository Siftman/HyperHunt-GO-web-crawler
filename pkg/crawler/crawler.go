package crawler

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"

	"github.com/yourusername/HyperHunt-GO-web-crawler/pkg/models"
)

func SitemapStatusCheck(baseURL string) (int, string) {
	status := 404
	url := baseURL

	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose)

	options := runner.Options{
		Methods:         "GET",
		InputTargetHost: goflags.StringSlice{baseURL + "/sitemap_index.xml", baseURL + "/sitemap.xml"},
		OnResult: func(r runner.Result) {
			if r.Err != nil {
				panic(r.Err)
			}
			fmt.Println(r.StatusCode, r.Input)
			if r.StatusCode == 200 {
				status = r.StatusCode
				url = r.Input
			}
		}}

	if err := options.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer httpxRunner.Close()

	httpxRunner.RunEnumeration()
	return status, url
}

func LinkSpider(baseURL string, unwantedPattern string) []string {
	var links []string
	c := colly.NewCollector(
		colly.AllowedDomains("shopino.app"),
		colly.MaxDepth(2),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("link found : ,", link)
		links = append(links, link)

		if unwantedPattern != "" {
			rej_regex, err := regexp.Compile(unwantedPattern)
			if err != nil {
				fmt.Println("Error compiling regex:", err)
				panic(err)
			}
			if !(rej_regex.MatchString(link)) {
				c.Visit(e.Request.AbsoluteURL(link))
			}
		} else {
			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	c.Visit(baseURL)
	return links
}

func Crawler(Webpage_link string) models.Product {
	first_collector := colly.NewCollector()
	second_collector := colly.NewCollector()
	var product models.Product
	var int_price int

	first_collector.OnHTML("script[type=\"application/ld+json\"]", func(e *colly.HTMLElement) {
		jsonStrings := strings.Split(e.Text, "\n")
		var product_schema models.Product_schema

		for _, jsonString := range jsonStrings {
			var genericMap map[string]interface{}
			if err := json.Unmarshal([]byte(jsonString), &genericMap); err != nil {
				fmt.Println("Error parsing JSON:", err)
				continue
			}
			if typeValue, ok := genericMap["@type"].(string); ok && typeValue == "Product" {
				if err := json.Unmarshal([]byte(jsonString), &product_schema); err != nil {
					fmt.Println(e.Request.URL)
					panic(err)
				} else {
					break
				}
			}
		}

		if product_schema.Name != "" && product_schema.Image != "" {
			switch v := product_schema.Image.(type) {
			case string:
				product.Image_Url = v
			case map[string]interface{}:
				if url, ok := v["url"].(string); ok {
					product.Image_Url = url
				}
			}
			product.Title = product_schema.Name

			switch v := product_schema.Offers.Price.(type) {
			case string:
				price, err := strconv.ParseFloat(v, 64)
				if err != nil {
					panic(err)
				} else {
					int_price = int(price)
				}
			case int:
				int_price = v
			}
			if product_schema.Offers.PriceCurrency == "IRR" {
				product.Price = int_price / 10
			} else {
				product.Price = int_price
			}
		} else {
			second_collector.Visit(Webpage_link)
		}
	})

	second_collector.OnHTML("head", func(e *colly.HTMLElement) {
		title := e.ChildAttr("meta[property=\"og:title\"]", "content")
		img_url := e.ChildAttr("meta[property=\"og:image\"]", "content")
		source_url := e.ChildAttr("meta[property=\"og:url\"]", "content")
		string_price := e.ChildAttr("meta[property=\"og:price:amount\"]", "content")

		if title != "" && img_url != "" {
			product.Title = title
			product.Image_Url = img_url
			product.Url = source_url

			if string_price != "" {
				price, err := strconv.ParseFloat(string_price, 64)
				if err != nil {
					panic(err)
				} else {
					product.Price = int(price)
				}
			}
		} else {
			fmt.Println(Webpage_link)
		}
	})

	first_collector.Visit(Webpage_link)
	return product
}

func SitemapScanner(Sitemap_url string) []string {
	var xml_list []string
	var url_list []string

	xml_collector := colly.NewCollector()
	url_collector := colly.NewCollector()

	xml_collector.OnXML("//loc", func(e *colly.XMLElement) {
		if strings.Split((e.Request.URL.Path), ".")[1] == "xml" {
			xml_list = append(xml_list, e.Text)
		}
	})

	url_collector.OnXML("//loc", func(e *colly.XMLElement) {
		url_list = append(url_list, e.Text)
	})

	xml_collector.Visit(Sitemap_url)

	for _, xml := range xml_list {
		url_collector.Visit(xml)
	}
	return url_list
}
