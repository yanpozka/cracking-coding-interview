// 1.1
import “strings”

func is_unique(str string) bool {
chars := make(map[string]bool)

for _, cc := range str {
	if _, contains := chars[cc]; contains {
	return false
} else {
	chars[cc] = true
}
}

return true

}

// 1.2
// C/C++
#include <strstr>

void reverse(char* str) {
	int L = strlen(str);
	if ( L <= 1 ) return;

	// “ab”   L = 2 ;  half = 1
	for (int ix = 0, half = L / 2; ix < half; ix++) {
char t = str[ix];
str[ix] = str[L-1-ix];
str[L-1-ix] = t;
}
}

// Go
func reverse(str string) string {
	len_str := len(str)
	new_str := make([]byte, len_str)
	for ix, k := len_str -1, 0; ix >= 0; ix++ {
		new_str[k] = str[ix]
		k++
}
return string(new_str)
}

// 1.3

func is_permutation(a string, b string) bool {
	if len(a) != len(b) {
	return false
}
a_map := make(map[string]int)
b_map := make(map[string]int)

for ix, L :=0, len(a); ix < L; ix++ {
	if _, contains := a_map[ix]; !contains {
	a_map[ix]
}
	}

}




