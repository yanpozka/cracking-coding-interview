// 2.4
package main

import "fmt"

type Node struct {
	next *Node
	data int
	// TODO: last *Node field to do O(1) aditions
}

func separate(head *Node, x int) *Node {
	var less, up *Node

	for c := head; c != nil; c = c.next {
		if c.data < x {
			less = add(less, c.data)
		} else if c.data >= x {
			up = add(up, c.data)
		}
	}

	var h *Node

	for c := less; c != nil; c = c.next {
		h = add(h, c.data)
	}
	for c := up; c != nil; c = c.next {
		h = add(h, c.data)
	}

	return h
}

func add(l *Node, v int) *Node {
	if l == nil {
		return &Node{data: v}
	}
	cursor := l
	for ; cursor.next != nil; cursor = cursor.next {
	}
	cursor.next = &Node{data: v}
	return l
}

func main() {
	var list *Node
	list = add(list, 6)
	list = add(list, 3)
	list = add(list, 2)
	list = add(list, 9)
	list = add(list, 5)

	printL(list)
	s := separate(list, 6)
	printL(s)
}

func printL(s *Node) {
	for cursor := s; cursor != nil; cursor = cursor.next {
		fmt.Print(cursor.data, " -> ")
	}
	fmt.Println("nil")
}
