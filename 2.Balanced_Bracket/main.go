package main

import "fmt"

func main() {
	fmt.Println(balanced_bracket("{[()]}"))       // Yes
	fmt.Println(balanced_bracket("{[(])}"))       // No
	fmt.Println(balanced_bracket("{(([])[])[]}")) // Yes
}

func balanced_bracket(s string) string {
	reverse := map[rune]rune{'}': '{', ']': '[', ')': '('}

	suite := []rune{}
	for _, bracket := range s {
		if bracket == '{' || bracket == '[' || bracket == '(' {
			suite = append(suite, bracket)
			continue
		}

		if bracket == '}' || bracket == ']' || bracket == ')' {
			if len(suite) == 0 || suite[len(suite)-1] != reverse[bracket] {
				return "NO"
			}

			suite = suite[:len(suite)-1]
		}
	}

	if len(suite) > 0 {
		return "NO"
	}
	return "YES"
}
