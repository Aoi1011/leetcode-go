package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test3Sum1(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	var output [][]int
	output = append(output, []int{-1, -1, 2})
	output = append(output, []int{-1, 0, 1})

	res := threeSum(nums)

	assert.Equal(t, output, res, "Pass")
}

func Test3Sum2(t *testing.T) {
	nums := []int{0, 1, 1}
	var output [][]int

	res := threeSum(nums)

	assert.Equal(t, output, res, "Pass")
}

func Test3Sum3(t *testing.T) {
	nums := []int{0, 0, 0}
	var output [][]int
	output = append(output, []int{0, 0, 0})

	res := threeSum(nums)

	assert.Equal(t, output, res, "Pass")
}
