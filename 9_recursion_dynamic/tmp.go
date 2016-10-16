// 9.1
// 9.2
// 9.3
// 9.4
//
// subsets (s) = []  			 ; if len(s) == 0
// subsets (s) = s[0] + subsets( s[1:] )  ; if len(s) > 0
//
func all_subsets(set []int) [][]int {
	result := make([][]int, 0)
	if len(set) == 0 {
	return result
}
result = append(result, set[0])
sub_result := all_subsets(set[1:])

	for ss := range sub_result {
		new_set := make([]int, len(ss) + 1)
		new_set[0] = set[0]
		new_set = append(new_set, ss)
		
		result = append(result, new_set)
}

result = append(result, new_set)
	
	return result 
}

// 9.5
func all_subsets(srt string) [][]int {
	set := []byte(srt)
	result := make([][]byte, 0)
	if len(set) == 0 {
		return result
}

len_set = len(set)






for ix := 0; ix < len_set; ix++ {
		as := all_subsets(append(set[:ix], set[ix+1:]...))
		
		for current_subset := range as {
			new_set := make([]byte, len(current_subset) + 1)
			new_set[0] = set[ix]
			copy(new_set[1:], current_subset)

			result = append(result, new_set) // !!
}
}
	
	return result 
}


