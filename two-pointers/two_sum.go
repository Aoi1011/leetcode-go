package twopointers

func twoSum(numbers []int, target int) []int {
	var res []int
	m := make(map[int]int)

	for index, num := range numbers {
		m[num] = index
	}

	for index, num := range numbers {
		diff := target - num
		if valIndex, found := m[diff]; found {
			res = append(res, index, valIndex)
		}
	}

	return res
}
