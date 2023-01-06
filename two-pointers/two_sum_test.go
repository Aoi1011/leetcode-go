package twopointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	output := []int{1, 2}

	res := twoSum(numbers, target)

	assert.Equal(t, output, res, "They should be equal")
}

func TestTwoSum1(t *testing.T) {
	numbers := []int{2, 3, 4}
	target := 6
	output := []int{1, 3}

	res := twoSum(numbers, target)

	assert.Equal(t, output, res, "They should be equal")
}

func TestTwoSum2(t *testing.T) {
	numbers := []int{-1, 0}
	target := -1
	output := []int{1, 2}

	res := twoSum(numbers, target)

	assert.Equal(t, output, res, "They should be equal")
}
