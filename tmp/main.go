package tmp

import (
	"fmt"
	"regexp"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func hello() string {
	return ""
}

type Item struct {
	Content string
}

var p = message.NewPrinter(language.SimplifiedChinese)

func test() {
}

func PrintAll() {
	//fmt.Println(string('\x77'))
	//fmt.Println(string('\u775e'))
	//fmt.Println(string('\U00100000'))

	str := "https://p3-zhuxiaobang-sign.shimolife.com/tos-cn-i-0000c2189/o8tIo9hB0BAfEmmEAIzAHABO8MHQ8yurzEanhC~tplv-hra1ikqiyn-rb:250:250:0:q80.webp?lk3s=993f2548&x-expires=1706586363&x-signature=5OVt9SaiA8T5cvhjAH4oFixHGeQ%3D"

	//runes := []rune(str)
	//bytes := []byte(str)

	//for _, r := range runes {
	//	fmt.Println(string(r))
	//}
	//
	//for _, b := range bytes {
	//	fmt.Printf("%x\n", b)
	//}

	reg := regexp.MustCompile("https?://([^.]+.[^.]+.[^/]+)")

	subMatches := reg.FindStringSubmatch(str)

	for _, match := range subMatches {
		fmt.Println(match)
	}
}

//func FormatFloat(value float64) string {
//	originVal := p.Sprintf("%.2f", value)
//	return strings.TrimRight(strings.TrimRight(originVal, "0"), ".")
//}
