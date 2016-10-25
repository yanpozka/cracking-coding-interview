// 1.1
package arrays_strings

import "testing"

func is_unique(str string) bool {

	chars := make(map[rune]bool)

	for _, cc := range str {
		if _, contains := chars[cc]; contains {
			return false
		} else {
			chars[cc] = true
		}
	}

	return true
}

func TestIsUnique(t *testing.T) {
	t.Log(is_unique("abcded"))
}
