package arrayshashing

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := [][]string{}
	output = append(output, []string{"bat"})
	output = append(output, []string{"nat", "tan"})
	output = append(output, []string{"ate", "eat", "tea"})

	res := GroupAnagrams(strs)

	if !reflect.DeepEqual(res, output) {
		t.Errorf("Err #1 %v, %v", res, output)
	}
}

func TestGroupAnagrams2(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := [][]string{}
	output = append(output, []string{"bat"})
	output = append(output, []string{"nat", "tan"})
	output = append(output, []string{"ate", "eat", "tea"})

	res := GroupAnagrams2(strs)

	if !reflect.DeepEqual(res, output) {
		t.Errorf("Err #1 %v, %v", res, output)
	}
}
