package main

import "fmt"

func ClosureFunc() func() {
	var val = 1
	ret := func() {
		val++
		fmt.Println(val)
	}
	val = 2
	return ret
}

func main() {
	func1 := ClosureFunc()
	func1()
	func1()

	rune := 'G'
	println(rune)
}
