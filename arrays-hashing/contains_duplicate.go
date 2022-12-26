package arrayshashing

func ContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		exist, found := m[num]
		if found || exist {
			return true
		}

		m[num] = true

	}

	return false
}
