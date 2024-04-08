package tmp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var subPageRegex = regexp.MustCompile("(http|https):\\/\\/www.feishu.cn\\/([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")

func test() {
	//subUrls, pageDTO := scanSubPage()
	//fmt.Printf("subUrls:%s, url:%s, title:%s", subUrls, pageDTO.Url, pageDTO.Title)

	urls := []string{"https://www.feishu.cn/hc/zh-CN/articles/694913918845"}
	scanPage(urls, 10)
}

type ToScanUrl struct {
	Url   string
	Depth int32
}

func scanPage(urls []string, depth int32) {
	scanMap := make(map[string]bool)
	scanMapLock := sync.Mutex{}

	ch := make(chan ToScanUrl, 100)
	for _, url := range urls {
		ch <- ToScanUrl{
			Url:   url,
			Depth: 1,
		}
	}

	timer := time.NewTicker(5 * time.Second)

	go func() {
		for {
			t := <-timer.C
			if len(ch) == 0 {
				fmt.Printf("t:%+v chan is empty,will close\n", t)
				close(ch)
				break
			}
		}
	}()

	for url := range ch {
		innerUrl := url
		go func() {
			if innerUrl.Depth > depth {
				return
			}
			subUrls, pageDTO := scanSubPage(innerUrl.Url)
			if pageDTO != nil {
				fmt.Printf("url:%s, title:%s, subUrls:%s\n", pageDTO.Url, pageDTO.Title, subUrls)
			}

			for _, subUrl := range subUrls {
				scanMapLock.Lock()
				if _, ok := scanMap[subUrl]; !ok {
					ch <- ToScanUrl{
						Url:   subUrl,
						Depth: innerUrl.Depth + 1,
					}
					scanMap[subUrl] = true
				}
				scanMapLock.Unlock()
			}
		}()
	}
}

type PageDTO struct {
	Url   string
	Title string
}

var titleRegex = regexp.MustCompile("<title>(.*)</title>")

func scanSubPage(url string) ([]string, *PageDTO) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	if len(sb) == 0 {
		return nil, nil
	}

	var pageDTO *PageDTO

	//fmt.Println(sb)
	subMatch := titleRegex.FindStringSubmatch(string(body))
	if len(subMatch) > 0 {
		pageDTO = &PageDTO{}
		pageDTO.Title = subMatch[1]
		pageDTO.Url = url
	}
	return subPageRegex.FindAllString(sb, 20), pageDTO
}
