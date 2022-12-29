package arrayshashing

import "sort"

// 1, 1, 1, 2, 2, 3
func topKFrequent(nums []int, k int) []int {
	var res []int
	firstM := make(map[int]int)
	secondM := make(map[int]int)

	// {1, 3}, {2, 2}, {3, 1}
	for _, num := range nums {
		firstM[num] += 1
	}

	//
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

func topKFrequent1(nums []int, k int) []int {
	m := make(map[int]int)
	max := 0

	for _, v := range nums {
		m[v] += 1
		if m[v] > max {
			max = m[v]
		}
	}

	mm := make(map[int][]int)
	for i, num := range m {
		mm[num] = append(mm[num], i)
	}

	ret := []int{}
	for i := max; i > 0; i-- {
		if elem, ok := mm[i]; ok {
			for _, v := range elem {
				ret = append(ret, v)
				if len(ret) == k {
					return ret
				}
			}
		}
	}

	return ret
}

func topKFrequent2(nums []int, k int) []int {
	m := make(map[int]int)
	max := 0

	for _, num := range nums {
		m[num] += 1
		if m[num] > max {
			max = m[num]
		}
	}

	mm := make(map[int][]int)
	for key, val := range m {
		mm[val] = append(mm[val], key)
	}

	ret := []int{}
	for i := max; i > 0; i-- {
		if elem, ok := mm[i]; ok {
			for _, v := range elem {
				ret = append(ret, v)
				if len(ret) == 2 {
					return ret
				}
			}
		}
	}

	return ret
}

// key, value = count
