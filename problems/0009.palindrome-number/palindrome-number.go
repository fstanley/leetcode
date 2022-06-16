package leetcode

import (
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	return s == rev(s)
}

func rev(s string) (result string) {
	for _, c := range s {
		result = string(c) + result
	}
	return result
}
