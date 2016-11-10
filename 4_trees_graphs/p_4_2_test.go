// 4.2
package main

import "container/list"

const (
	N     = 1000
	empty = -1
)

func existRoute(Graph [N][N]int, x, y int) bool {
	var visited [N]bool
	visited[0] = true
	queue := list.New()
	queue.PushBack(x)

	for queue.Len() > 0 {
		node := queue.First()
		current_vertex := node.Value.(int)
		queue.Remove(node)

		for ix := 0; ix < N; ix++ {
			if Graph[current_vertex][ix] != empty && !visited[ix] {
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
