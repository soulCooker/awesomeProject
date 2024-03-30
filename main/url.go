package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	params := url.Values{}
	params.Add("aaa", "123")
	params.Add("bbbb", "452")

	fmt.Println(fmt.Sprintf(params.Encode()))

	values := map[string]string{"val": fmt.Sprintf(params.Encode())}

	strBytes, _ := json.Marshal(values)
	fmt.Println(string(strBytes))
}
