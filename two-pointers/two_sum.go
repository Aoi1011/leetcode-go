package twopointers

func twoSum(numbers []int, target int) []int {
	var res []int
	m := make(map[int]int)

	for index, num := range numbers {
		m[num] = index + 1
	}

	for index, num := range numbers {
		diff := target - num
		if valIndex, found := m[diff]; found {
			res = append(res, index+1, valIndex)
		}

		if len(res) == 2 {
			break
		}
	}

	return res
}

func TwoSum1(numbers []int, target int) []int {
	var res []int
	left := 0
	right := len(numbers) - 1

	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			res = append(res, left+1, right+1)
			break
		} else if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return res
}
