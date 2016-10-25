// 1.3
package arrays_strings

import "testing"

func is_permutation(a string, b string) bool {

	if len(a) != len(b) {
		return false
	}

	// a_map := make(map[string]int)
	// b_map := make(map[string]int)

	// for ix, L := 0, len(a); ix < L; ix++ {
	// 	if _, contains := a_map[ix]; !contains {
	// 		a_map[ix]
	// 	}
	// }

	// TODO: stuff

	return true
}

func TestIsPermutation(t *testing.T) {
	t.Log(is_permutation("abc", "bcd"))
}
