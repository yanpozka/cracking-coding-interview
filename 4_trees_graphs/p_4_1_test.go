// 4.1
package main

import "math"

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	data  int
}

func height(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return 1 + math.Max(float64(height(root.left)), float64(height(root.right)))
}

func isBalanced(root *BinaryTree) bool {
	if root == nil {
		return true
	}

	return math.Abs(float64(height(root.left)-height(root.right))) < 1
}
