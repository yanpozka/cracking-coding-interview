type BinaryTree struct {
	left *BinaryTree
	right *BinaryTree
	data int
}

// 4.1
import “math”

func height(root *BinaryTree) int {
	if root == nil {
		return 0
}
return 1 + math.Max( float64(height(root->left)), float64(height(root->right)) )
}

func isBalanced(root *BinaryTree) bool {
	if root == nil {
		return true
}
return math.Abs(float64(height(root->left) - height(root->right))) < 1
}

// 4.2
import “container/list”

const N = 1000

var Graph [N][N]int

func existRoute(x, y  int) bool {
var visited [N]bool
visited[0] = true
	queue := list.New()
	queue.PushBack(x)

	for queue.Len() > 0 {
		node := queue.First()
		current_vertex := node.Value.(int)
		queue.Remove(node)

	

	for ix := 0; ix < N; ix++ {
		if  Graph[current_vertex][ix] != -1 && !visited[ix] {
			if ix == y {
				return true
}
	visited[ix] = true
queue[queueCount] = ix
queueCount++
}
}
}
return false
}

