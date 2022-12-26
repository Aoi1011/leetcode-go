package arrayshashing

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	contains := ContainsDuplicate(nums)

	if !contains {
		t.Errorf("Error #1 %t", contains)
	}
}
