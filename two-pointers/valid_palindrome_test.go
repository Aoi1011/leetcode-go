package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	output := true

	res := IsPalindrome(s)

	if !res {
		t.Errorf("Fail: expected %t, actual: %t", output, res)
	}
}
