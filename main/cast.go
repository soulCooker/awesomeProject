package main

import (
	"fmt"

	"github.com/spf13/cast"
)

func main() {
	var str *string
	val := cast.ToInt64(str)
	fmt.Println(val)
}
