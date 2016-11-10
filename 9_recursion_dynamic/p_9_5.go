// 9.5
package main

import (
	"container/list"
	"fmt"
)

func permutations(srt string) []*list.List {
	result := make([]*list.List, 0)
	if len(srt) == 0 {
		return result
	}

	// TODO: everything

	return result
}

func main() {
	fmt.Println(permutations("abc"))
}
