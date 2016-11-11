// 2.1 Remove duplicates
package main

import "fmt"

type Node struct {
	next *Node
	data int
}

func pushBack(head *Node, d int) {
	if head == nil {
		head = &Node{data: head.data}
		return
	}
	cursor := head
	for {
		if cursor.next == nil {
			break
		}
		cursor = cursor.next
	}
	cursor.next = &Node{data: head.data}
}

func removeDuplicates(head *Node) *Node {
	if head == nil {
		return nil
	}
	new_list := &Node{data: head.data}
	nc := new_list

	for cursor := head.next; cursor != nil; cursor = cursor.next {
		if !contains(new_list, cursor.data) {
			nc.next = &Node{data: cursor.data}
			nc = nc.next
		}
	}
	return new_list
}

func contains(head *Node, d int) bool {
	var found bool
	for cursor := head; cursor != nil && !found; cursor = cursor.next {
		if cursor.data == d {
			found = true
			break
		}
	}
	return found
}

func main() {

	list := &Node{
		data: 2,
		next: &Node{
			data: 3,
			next: &Node{
				data: 2,
				next: &Node{
					data: 9,
					next: &Node{data: 3},
				},
			},
		},
	}
	printL(list)

	s := removeDuplicates(list)

	printL(s)
}

func printL(s *Node) {
	for cursor := s; cursor != nil; cursor = cursor.next {
		fmt.Print(cursor.data, " -> ")
	}
	fmt.Println("nil")
}
