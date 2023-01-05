package twopointers

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	output := true

	res := IsPalindrome(s)

	if res != output {
		t.Errorf("Fail1: expected %t, actual: %t", output, res)
	}

	s = "race a car"
	output = false

	res = IsPalindrome(s)

	if res != output {
		t.Errorf("Fail1: expected %t, actual: %t", output, res)
	}

	s = " "
	output = true

	res = IsPalindrome(s)

	if res != output {
		t.Errorf("Fail1: expected %t, actual: %t", output, res)
	}
}
