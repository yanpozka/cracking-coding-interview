// 4.1
package main

import "math"

type binaryTree struct {
	left  *binaryTree
	right *binaryTree
	data  int
}

func height(root *binaryTree) int {
	if root == nil {
		return 0
	}
	return 1 + math.Max(float64(height(root.left)), float64(height(root.right)))
}

func isBalanced(root binaryTree) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(height(root.left)-height(root.right))) < 1
}
