package main

import "fmt"

func main() {
	hanoi(7, "a", "b", "c")
}

func hanoi(n int, a, b, c string) {
	if n == 1 {
		fmt.Println(a, "->", c)
		return
	}
	hanoi(n-1, a, c, b)
	fmt.Println(a, "->", c)
	hanoi(n-1, b, a, c)
}
