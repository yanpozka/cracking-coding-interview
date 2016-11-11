// 9.6
package main

import (
	"fmt"
)

// https://play.golang.org/p/yIECWdEN8h
//
func main() {
	all := [][]string{[]string{""}, []string{"()"}} //, []string{"(())", "()()"}}

	for ix := 2; ix < 9; ix++ {
		var cs []string

		for k, val := range all[ix-1] {
			cs = append(cs, "("+val+")")
			cs = append(cs, "()"+val)
			if k < len(all[ix-1])-1 {
				cs = append(cs, val+"()")
			}
		}

		all = append(all, cs)
	}

	for _, s := range all {
		fmt.Println(s)
	}
}
