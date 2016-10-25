// 1.2
package arrays_strings

import "testing"

// // C/C++
// #include <strstr>

// void reverse(char* str) {
// 	 int L = strlen(str);
// 	 if ( L <= 1 ) return;

// 	 // “ab”   L = 2 ;  half = 1
// 	 for (int ix = 0, half = L / 2; ix < half; ix++) {
//  	char t = str[ix];
// 		str[ix] = str[L-1-ix];
// 		str[L-1-ix] = t;
//   }
// }

// Go
func reverse(str string) string {
	len_str := len(str)
	new_str := make([]byte, len_str)

	for ix, k := len_str-1, 0; ix >= 0; ix-- {
		new_str[k] = str[ix]
		k++
	}

	return string(new_str)
}

func TestReverseStr(t *testing.T) {
	t.Log(reverse("abcdef"))
}
