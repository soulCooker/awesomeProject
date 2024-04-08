package tmp

import "fmt"

// )(())(())

func test() {
	//
	count := findLongestBracketString("()((()))((())")
	fmt.Println(count)
}

type Counter struct {
	ch  rune
	len int
}

func findLongestBracketString(str string) int {
	longestLen := 0
	curLen := 0

	stack := make([]*Counter, 0)

	for _, ch := range str {
		if ch == ')' {
			if len(stack) > 0 {
				if stack[len(stack)-1].ch == '(' {
					curLen = stack[len(stack)-1].len + 2
					stack = stack[0 : len(stack)-1]
					if curLen > longestLen {
						longestLen = curLen
					}
				}
			}

			if len(stack) > 0 {
				stack[len(stack)-1].len += curLen
			}
		} else if ch == '(' {
			stack = append(stack, &Counter{
				ch:  ch,
				len: 0,
			})
		}

	}
	return longestLen
}
