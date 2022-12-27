package arrayshashing

func TwoSum(nums []int, target int) []int {
	var arr []int

	for index, num := range nums {
		for i := index + 1; i < len(nums); i++ {
			if target-num == nums[i] {
				return append(arr, index, i)
			}
		}
	}

	return arr
}
