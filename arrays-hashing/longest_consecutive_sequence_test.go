package arrayshashing

import "testing"

func TestLongestConsecutiveSequence(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	output := 4

	res := longestConsecutive3(nums)

	if res != output {
		t.Errorf("Fail: expected %v, actual %v", output, res)
	}

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	output = 9

	res = longestConsecutive3(nums)

	if res != output {
		t.Errorf("Fail: expected %v, actual %v", output, res)
	}
}
