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

func GroupAnagrams2(strs []string) [][]string {
	mp := make(map[[26]int][]string)

	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}

	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}

	return res
}
