package arrayshashing

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	var maxes []int
	max := 1
	// var m map[int]int
	lenNums := len(nums)
	// start := 0

	sort.Ints(nums)

	for i := 0; i < lenNums; i++ {
		if i == lenNums-1 {
			continue
		}

		if nums[i+1] != nums[i]+1 {
			maxes = append(maxes, max)
			max = 1
			continue
		}

		max += 1

	}

	sort.Ints(maxes)

	fmt.Printf("Maxes, %v", maxes)

	return maxes[len(maxes)-1]
}

func longestConsecutive2(nums []int) int {
	mapper := make(map[int]int)
	for _, v := range nums {
		mapper[v] = 1
	}
	ans := 0
	for _, v := range nums {
		if mapper[v-1] == 1 {
			continue
		}
		y := v + 1
		for mapper[y] == 1 {
			y++
		}
		if y-v > ans {
			ans = y - v
		}
	}

	return ans
}
