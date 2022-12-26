package arrayshashing

import "testing"

func TestIsAnagram(m *testing.T) {
	s := "anagram"
	t := "nagaram"

	match := IsAnagram(s, t)
	if !match {
		m.Errorf("Error #1 %t", match)
	}

	s = "rat"
	t = "car"

	match = IsAnagram3(s, t)
	if match {
		m.Errorf("Error #1 %t", match)
	}
}
