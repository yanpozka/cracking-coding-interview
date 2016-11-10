package treesGraphs

import "testing"

type binaryTree struct {
	left  *binaryTree
	right *binaryTree
	data  int
}

func makeTree(nums []int) *binaryTree {
	ln := len(nums)

	if ln == 0 {
		return nil
	}
	if ln == 1 {
		return &binaryTree{data: nums[0]}
	}

	half := ln / 2

	return &binaryTree{
		data:  nums[half],
		left:  makeTree(nums[0:half]),
		right: makeTree(nums[half+1:]),
	}
}

func inOrder(r *binaryTree) []int {
	var result []int
	if r == nil {
		return result
	}

	result = append(result, inOrder(r.left)...)
	result = append(result, r.data)
	return append(result, inOrder(r.right)...)
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// https://play.golang.org/p/mz_ImRDf_i
//
func TestMakeTree(t *testing.T) {
	a := []int{1, 3, 4, 5, 6, 7, 8, 11, 13, 14, 15, 16, 20, 24}
	b := inOrder(makeTree(a))

	if testEq(ns, r) == false {
		t.Fatalf("Non equals slices: %v and %v \n", a, b)
	}
}
