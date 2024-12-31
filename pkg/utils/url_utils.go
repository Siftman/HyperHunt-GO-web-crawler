package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

func Proper_link_detector(Urls []string) []string {
	keywords := []string{"products", "product", "item", "buy", "shop", "detail", "catalog"}
	unwantedKeywords := []string{"blog"}

	accepted_urls := []string{}

	keywordsPattern := `\b(` + strings.Join(keywords, "|") + `)\b`
	unwantedPattern := `\b(` + strings.Join(unwantedKeywords, "|") + `)\b`

	acc_regex, err := regexp.Compile(keywordsPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic(err)
	}

	rej_regex, err := regexp.Compile(unwantedPattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic(err)
	}

	for _, url := range Urls {
		if acc_regex.MatchString(url) {
			if rej_regex.MatchString(url) {
				fmt.Printf("Special No match: %s\n", url)
				continue
			} else {
				accepted_urls = append(accepted_urls, url)
				fmt.Printf("Match found: %s\n", url)
			}
		} else {
			fmt.Printf("No match: %s\n", url)
		}
	}
	return accepted_urls
}

func Frequency_finder(urls []string) string {
	var mostFrequent string
	var maxCount int

	frequency := make(map[string]int)

	for _, str := range urls {
		u, err := url.Parse(str)
		if err != nil {
			panic(err)
		}
		first_token := strings.Split(u.Path, "/")
		fmt.Println(first_token)
		frequency[first_token[1]]++
	}

	for str, count := range frequency {
		if count > maxCount {
			maxCount = count
			mostFrequent = str
		}
	}

	return mostFrequent
}

func Filter_csv(urls []string, freq_path string) []string {
	var Final_links []string

	for _, str := range urls {
		u, err := url.Parse(str)
		if err != nil {
			panic(err)
		}
		first_token := strings.Split(u.Path, "/")
		if first_token[1] == freq_path {
			Final_links = append(Final_links, str)
		}
	}
	return Final_links
}
