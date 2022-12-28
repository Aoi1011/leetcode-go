package arrayshashing

import "sort"

func topKFrequent(nums []int, k int) []int {
	var res []int
	firstM := make(map[int]int)
	secondM := make(map[int]int)

	for _, num := range nums {
		firstM[num] += 1
	}

	var temp []int
	for k, v := range firstM {
		temp = append(temp, v)
		secondM[v] = k
	}
	sort.Ints(temp)

	for i := len(temp) - 1; i >= len(temp)-k; i-- {
		sec := secondM[temp[i]]
		res = append(res, sec)
		secondM[temp[i]] = -1
	}
	return res
}
