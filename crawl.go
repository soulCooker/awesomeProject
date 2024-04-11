package main

import (
	"fmt"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func Crawl(url string, depth int) {
	ch := make(chan []string)
	currentLevel := []string{url}
	errCh := make(chan string)
	pattern := regexp.MustCompile("/talks/")
	discovered := make(map[string]interface{})

	for len(currentLevel) > 0 && depth > 0 {
		depth--
		nextLevel := make([]string, 0)
		for _, s := range currentLevel {
			go doCrawl(s, ch, errCh, func(c string) {
				_, _ = fmt.Fprintln(os.Stdout, c)
			}, pattern.MatchString)
		}

	label:
		for i := 0; i < len(currentLevel); i++ {
			timeOut := time.After(2 * time.Second)
			select {
			case nextLinks := <-ch:
				for _, link := range nextLinks {
					shortLink := link
					if sharpIndex := strings.Index(link, "#"); sharpIndex > 0 {
						shortLink = link[0:strings.Index(link, "#")]
					}

					if _, ok := discovered[shortLink]; ok {
						continue
					} else {
						nextLevel = append(nextLevel, shortLink)
						discovered[shortLink] = 1
					}
				}
			case errCh := <-errCh:
				fmt.Println(errCh)
			case curTime := <-timeOut:
				fmt.Printf("timeout:%s", curTime)
				break label
			default:
				time.Sleep(time.Millisecond * 50)
				i--
			}
		}
		currentLevel = nextLevel
	}
}

func doCrawl(url string, c chan []string, errChan chan string, handler func(title string), filter func(url string) bool) {
	//start := time.Now()
	//defer func() {
	//	fmt.Printf("Url:%s, Cost:%v\n", url, time.Now().Sub(start))
	//}()
	resp, err := http.Get(url)
	if err != nil {
		errChan <- "network err url: " + url
		return
	}
	node, err := html.Parse(resp.Body)
	if err != nil {
		errChan <- "parse err, url:" + url
		return
	}

	title, ok := parseTitle(node)
	if ok {
		handler(fmt.Sprintf("title:%s, url:%s", title, url))
	}

	rootUrl, _ := url2.Parse(url)
	urls := parseSubLink(node, filter)
	for i, subUrl := range urls {
		rootDir := fmt.Sprintf("%s://%s", rootUrl.Scheme, rootUrl.Host)
		if strings.HasPrefix(subUrl, "/") {
			urls[i] = fmt.Sprintf("%s%s", rootDir, subUrl)
		}
	}
	c <- urls
}

func parseTitle(node *html.Node) (string, bool) {
	if isTitleNode(node) {
		return strings.TrimSpace(node.FirstChild.Data), true
	}
	for node = node.FirstChild; node != nil; node = node.NextSibling {
		title, ok := parseTitle(node)
		if ok {
			return strings.TrimSpace(title), true
		}
	}
	return "", false
}

func isTitleNode(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "title"
}

func isHyperLink(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

func parseSubLink(node *html.Node, matcher func(string) bool) []string {
	res := make([]string, 0)

	if isHyperLink(node) {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" && matcher(attribute.Val) {
				res = append(res, attribute.Val)
				break
			}
		}
	}
	for node = node.FirstChild; node != nil; node = node.NextSibling {
		subLinks := parseSubLink(node, matcher)
		res = append(res, subLinks...)
	}
	return res
}
