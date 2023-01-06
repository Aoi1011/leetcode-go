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
