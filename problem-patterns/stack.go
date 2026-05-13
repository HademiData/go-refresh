package problempatterns

import "strings"

func revString(s string) string {

	stack := []rune{}

	for _, ch := range s {
		stack = append(stack, ch)
	}

	ans := ""

	for len(stack) > 0 {

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		ans += string(top)
	}

	return ans
}

func validParenthesis(s string) bool {

	stack := []rune{}

	pairs := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}

	openingBrackets := map[rune]bool{'[': true, '{': true, '(': true}

	for _, ch := range s {

		if strings.ContainsRune("[{()}]", ch) {

			if openingBrackets[ch] {
				stack = append(stack, ch)
			} else {

				if len(stack) == 0 {
					return false
				}

				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pairs[ch] != top {
					return false
				}

			}
		}
	}

	return len(stack) == 0
}

func warmerTemp(temp []int) []int {

	ans := make([]int, len(temp))

	stack := []int{}

	for i := 0; i < len(temp); i++ {

		for len(stack) > 0 && temp[i] > temp[stack[len(stack)-1]] {

			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ans[idx] = i - idx

		}

		stack = append(stack, i)
	}

	return ans
}
