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
	total := 1
	zeroCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount += 1
			continue
		}
		total *= nums[i]
	}

	for i := 0; i < len(nums); i++ {
		if zeroCount == 0 {
			res = append(res, total/nums[i])
		} else if zeroCount == 1 && nums[i] == 0 {
			res = append(res, total)
		} else {
			res = append(res, 0)
		}
	}

	return res
}
