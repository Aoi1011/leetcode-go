package arrayshashing

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	output := []int{24, 12, 8, 6}

	res := productExceptSelf(nums)

	if reflect.DeepEqual(res, output) {
		t.Errorf("Error: %v, %v", output, res)
	}

	nums = []int{-1, 1, 0, -3, 3}
	output = []int{0, 0, 9, 0, 0}

	res = productExceptSelf(nums)

	if reflect.DeepEqual(output, res) {
		t.Errorf("Error: %v, %v", output, res)
	}
}
