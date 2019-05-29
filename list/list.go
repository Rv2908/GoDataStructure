package list

import "fmt"

// A linked list is a linear collection of data elements, 
// in which linear order is not given by their physical placement in memory. 
// Each pointing to the next node by means of a pointer. 
// It is a data structure consisting of a group of nodes which together represent a sequence. 
// Here is source code of the Go Program to Implement Single Unsorted Linked List
// func LinkedList()

// First Retrive the first element of the linked list
func (L *List) First() *Node {
	return L.head
}
// Last Retrive the last element of the linked list
func (L *List) Last() *Node {
	return L.tail
}

// Next Retrive the next element in the node
func (N *Node) Next() *Node {
	return N.next
}
// Prev Retrive the prev element in the node
func (N *Node) Prev() *Node {
	return N.prev
}

// Push This will push the element in the list
func (L *List) Push(key interface{}) {
	node := &Node{
		key: key,
	}
	// if the head node is nil then we create a head node
	if L.head == nil {
		L.head = node
	} else {
		// Now assign the value to the next node of the next node of tail
		L.tail.next = node
		node.prev = L.tail
	}

	// Also assign the current node to the tail
	L.tail = node
}


// IterateForward it prints linked list from head to tail i.e. start to end 
func(L *List) IterateForward(){
	node := L.First()
	for {
		fmt.Println(node.key)
		node = node.Next()
		if node == nil{
			break
		}
	}
}

// IterateBackward it iterate Linked List from tail to head
func(L *List) IterateBackward(){

	node := L.Last()
	for {
		fmt.Println(node.key)
		node = node.Prev()
		if node == nil{
			break
		}
	}
}

// func (L *List) Reverse() {
// 	curr := L.First()
// 	var prev *Node
// 	L.tail = L.head

// 	for curr!=nil{
// 		next := curr.next
// 		curr.next = prev
// 		prev= curr
// 		curr= next
// 	}
// 	L.head= prev
// 	L.IterateList()
// }
