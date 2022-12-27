package arrayshashing

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := []int{0, 1}

	res := TwoSum(nums, target)

	if !reflect.DeepEqual(ans, res) {
		t.Errorf("Err #1: %v, %v", ans, res)
	}

	nums = []int{3, 2, 4}
	target = 6
	ans = []int{1, 2}

	res = TwoSum(nums, target)

	if !reflect.DeepEqual(ans, res) {
		t.Errorf("Err #1: %v, %v", ans, res)
	}

}
