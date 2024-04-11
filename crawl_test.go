package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCrawl(t *testing.T) {
	start := time.Now()
	Crawl("https://go.dev/talks/", 5)
	fmt.Printf("cost:%v\n", time.Now().Sub(start))
}
