package list

// Node Contain the nodes of the list prev and next
type Node struct{
	prev *Node
	next *Node
	key interface{}
}

// List  Actual Linked List
type List struct{
	head *Node
	tail *Node
}

