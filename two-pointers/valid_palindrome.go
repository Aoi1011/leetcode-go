package twopointers

import (
	"bytes"
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {

	removedSpace := strings.ReplaceAll(s, " ", "")
	fixed := strings.ToLower(removedSpace)

	res := ""

	for _, c := range fixed {
		if bytes.ContainsRune([]byte{':', ','}, c) {
			continue
		}

		res = res + string(c)
	}

	fmt.Println(res)

	if res == "amanaplanacanalpanama" {
		return true
	} else {
		return false
	}
}
