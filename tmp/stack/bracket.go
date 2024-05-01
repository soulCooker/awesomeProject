package stack

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if v == ')' && top != '(' {
				return false
			} else if v == ']' && top != '[' {
				return false
			} else if v == '}' && top != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
