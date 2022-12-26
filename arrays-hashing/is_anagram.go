package arrayshashing

import "reflect"

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make([]int, 26)

	for index, _ := range count {
		count[index] = 0
	}

	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
		count[int(t[i])-int('a')]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}

	return true
}

func IsAnagram2(s string, t string) bool {
	sMap := map[byte]int{}
	tMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sMap[s[i]] = sMap[s[i]] + 1
	}

	for i := 0; i < len(t); i++ {
		tMap[t[i]] = tMap[t[i]] + 1
	}

	return reflect.DeepEqual(sMap, tMap)
}

func IsAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[int]int)
	for i := range s {
		m[int(s[i])]++
	}

	for i := range t {
		m[int(t[i])]--
		if m[int(t[i])] < 0 {
			return false
		}
	}

	return true
}
