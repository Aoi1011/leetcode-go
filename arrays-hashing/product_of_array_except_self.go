package arrayshashing

// time limit exceeded
func productExceptSelf(nums []int) []int {
	var res []int

	for index := range nums {
		total := 1
		for i := 0; i < len(nums); i++ {
			if i == index {
				continue
			}

			total *= nums[i]
		}

		res = append(res, total)
	}

	return res
}

// solution 1
func productExceptSelf1(nums []int) []int {
	var res []int

	return res
}
