package twopointers

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	n := len(nums)

	sort.Ints(nums)

	for num1Idx := 0; num1Idx < n-2; num1Idx++ {
		if num1Idx > 0 && nums[num1Idx] == nums[num1Idx-1] {
			continue
		}

		num2Idx := num1Idx + 1
		num3Idx := n - 1

		for num2Idx < num3Idx {
			sum := nums[num1Idx] + nums[num2Idx] + nums[num3Idx]
			if sum == 0 {
				res = append(res, []int{nums[num1Idx], nums[num2Idx], nums[num3Idx]})

				num3Idx--

				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx+1] {
					num3Idx--
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
		}
	}

	return res
}

func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	for index, val := range nums {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		left := index + 1
		right := len(nums) - 1

		for left < right {
			threeSum := val + nums[left] + nums[right]
			if threeSum == 0 {
				res = append(res, []int{val, nums[left], nums[right]})

				// [-2, -2, 0, 0, 2, 2]
				left++
				for nums[left] == nums[left-1] && left < right {
					left++
				}
			} else if threeSum > 0 {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
