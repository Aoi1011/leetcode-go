package arrayshashing

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	strSli := strings.Split(str, "")
	sort.Strings(strSli)
	return strings.Join(strSli, "")
}

func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}
	indexMap := make(map[int]bool)

	for index, str := range strs {
		var temArr []string

		if _, found := indexMap[index]; found {
			continue
		}

		temArr = append(temArr, str)

		for j := index + 1; j < len(strs); j++ {
			if sortString(str) == sortString(strs[j]) {
				temArr = append(temArr, strs[j])
				indexMap[j] = true
			}
		}

		res = append(res, temArr)
	}

	return res
}
