package main

import "fmt"

func main() {
	fmt.Println(find_highest_palindrome("3943", 1))   // Output: 3993
	fmt.Println(find_highest_palindrome("932239", 2)) // Output: 992299
}

func find_highest_palindrome(s string, k int) string {
	changed := make([]bool, len(s))
	return highest_palindrome(s, k, 0, len(s)-1, changed)
}

func highest_palindrome(s string, k int, left int, right int, changed []bool) string {
	if left >= right {
		if k >= 0 {
			return s
		} else {
			return "-1"
		}
	}

	if s[left] == s[right] {
		return highest_palindrome(s, k, left+1, right-1, changed)
	}

	if k <= 0 {
		return "-1"
	}

	if k >= 2 || (k == 1 && changed[left] || changed[right]) {
		if s[left] != '9' {
			newS := s[:left] + "9" + s[left+1:]
			changed[left] = true
			if res := highest_palindrome(newS, k-1, left+1, right-1, changed); res != "-1" {
				return res
			}
			changed[left] = false
		}

		if s[right] != '9' {
			newS := s[:right] + "9" + s[right+1:]
			changed[right] = true
			if res := highest_palindrome(newS, k-1, left+1, right-1, changed); res != "-1" {
				return res
			}
			changed[right] = false
		}
	}

	if s[left] != s[right] {
		maxChar := s[left]
		if s[right] > maxChar {
			maxChar = s[right]
		}
		newS := s[:left] + string(maxChar) + s[left+1:right] + string(maxChar) + s[right+1:]
		changed[left], changed[right] = true, true
		return highest_palindrome(newS, k-1, left+1, right-1, changed)
	}

	return "-1"
}
