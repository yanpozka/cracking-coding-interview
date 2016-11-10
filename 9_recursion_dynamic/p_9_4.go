// 9.4
package main

import (
	"container/list"
	"fmt"
	"math"
)

//
// subsets (s) = []  ; if len(s) == 0
// subsets (s) = [ s[0] ] + for each subsets( s[1:] )  ; if len(s) > 0
//
func all_subsets(set []int) []*list.List {
	result := make([]*list.List, 0)

	if len(set) == 0 {
		return result
	}
	l := list.New()
	l.PushBack(set[0])
	result = append(result, l)

	osets := all_subsets(set[1:])

	for _, lo := range osets {
		new_list := list.New()
		new_list.PushBack(set[0])
		new_list.PushBackList(lo)

		result = append(result, new_list)
	}

	for _, lo := range osets {
		result = append(result, lo)
	}

	return result
}

func main() {
	set := []int{1, 2, 3, 4, 5}
	ss := all_subsets(set)

	fmt.Println(len(ss), math.Pow(2, float64(len(set)))-1)

	for _, set := range ss {
		for e := set.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v ", e.Value)
		}
		fmt.Println()
	}
}
