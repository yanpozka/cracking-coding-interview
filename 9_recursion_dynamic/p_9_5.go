// 9.5
package main

import (
	"container/list"
	"fmt"
)

// "abc"
// "acb"
//
func permutations(str string) *list.List {
	result := list.New()

	if str == "" {
		return result
	}

	for ix, c := range str {
		ol := permutations(str[:ix] + str[ix+1:])

		for e := ol.Front(); e != nil; e = e.Next() {
			result.PushBack(string(c) + e.Value.(string))
		}

		if ol.Len() == 0 {
			result.PushBack(string(c))
		}
	}

	return result
}

// https://play.golang.org/p/t2qmo99_N_
//
func main() {
	str := "abc"
	l := permutations(str)

	fmt.Println(l.Len(), "==", fac(len(str)))

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v\n", e.Value)
	}
}

func fac(n int) uint64 {
	if n <= 0 || n > len(_f) {
		return 0
	}
	return _f[n]
}

var _f = [...]uint64{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600, 6227020800,
	87178291200, 1307674368000, 20922789888000, 355687428096000, 6402373705728000, 121645100408832000, 2432902008176640000,
	14197454024290336768, 17196083355034583040, 8128291617894825984, 10611558092380307456, 7034535277573963776, 16877220553537093632,
	12963097176472289280, 12478583540742619136, 11390785281054474240, 9682165104862298112, 4999213071378415616, 12400865694432886784,
	3400198294675128320, 4926277576697053184, 6399018521010896896, 9003737871877668864, 1096907932701818880, 4789013295250014208, 2304077777655037952, 18376134811363311616}
