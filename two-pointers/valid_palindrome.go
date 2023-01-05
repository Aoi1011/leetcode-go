package twopointers

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	removedSpace := strings.ReplaceAll(s, " ", "")
	lower := strings.ToLower(removedSpace)
	fixed := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(lower, "")

	for i := 0; i < len(fixed)/2; i++ {
		if fixed[i] != fixed[len(fixed)-i-1] {
			return false
		}
	}

	return true
}
