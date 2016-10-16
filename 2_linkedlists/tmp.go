// 2.1 Remove duplicates
type Node struct {
	next *Node
	data int
}

func pushBack(head *Node, d int) {
	if head == nil {
	head = &Node{data: head->data}
} else {
cursor := head
for {
	if cursor->next == nil {
		break
}
cursor = cursor->next
}
	cursor->next = &Node{data: head->data}
}
}

func removeDuplicates(head *Node) *Node {
	if head == nil {
		return
}
	new_list := &Node{data: head->data}
	
	for cursor := head->next; cursor != nil; cursor = cursor->next {
		if !contains(new_list, cursor->data) {
			add(new_list, cursor->data)
}
}
return new_list
}

func contains(head *Node, d int) bool {
	var found bool
for cursor := head; cursor != nil && !found; cursor = cursor->next {
		if cursor->data == d {
			found = true
}
}
return found
}

// 2.2

// 2.3

// 2.4
func separte(head *Node, x int) *Node {
	min_node, max_node := nil, nil
	
	for cursor := head; cursor != nil; cursor = cursor->next {
		if  cursor->data <= x {
			add(min_node, x)
} else {
	add(max_node, x)
}
}
if min_cursor != nil {
min_cursor := head
for {
	if cursor->next == nil {
		break
}
cursor = cursor->next
}
min_cursor->next = max_node
return min_cursor
}
return max_node
}

