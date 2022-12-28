package arrayshashing

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	ans := []int{1, 2}

	res := topKFrequent(nums, k)

	if !reflect.DeepEqual(ans, res) {
		t.Errorf("Err #1: %v, %v", ans, res)
	}

	// nums = []int{3, 2, 4}
	// target = 6
	// ans = []int{1, 2}

	// res = TwoSum(nums, target)

	// if !reflect.DeepEqual(ans, res) {
	// 	t.Errorf("Err #1: %v, %v", ans, res)
	// }

}
