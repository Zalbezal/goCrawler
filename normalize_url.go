package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)

	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil

	/*
		Original implementation
		// if len(url) == 0 {
		// 	return "", errors.New("please insert a url")
		// }
		// urlArray := strings.Split(url, "//")
		// formattedString := urlArray[1]

		// if string(formattedString[len(formattedString)-1]) == "/" {
		// 	formattedString = formattedString[:len(formattedString)-1]
		// }
		// return formattedString, nil
	*/
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	urls := []string{}
	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, err
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					urls = append(urls, a.Key)
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}

		}
		f(doc)
	}
	return urls, nil
}
