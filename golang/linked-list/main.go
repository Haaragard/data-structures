package main

import "fmt"

func main() {
	linkedList := &LinkedList{}
	linkedList.InsertAtEnd("First node")
	linkedList.InsertAtEnd("Second node")
	linkedList.InsertAtEnd("Third node")
	linkedList.InsertAtEnd("Fourth node")
	linkedList.InsertAtEnd("Fifth node")
	linkedList.InsertAtEnd("Sixth node")

	linkedList.PrintList()

	fmt.Printf("This list has %d items\n", linkedList.FindLength())
}

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	root     *Node
	lastNode *Node
}

func (self *LinkedList) InsertAtEnd(data string) {
	node := &Node{
		data: data,
		next: nil,
	}

	if self.root == nil {
		self.root = node
	} else {
		self.lastNode.next = node
	}

	self.lastNode = node
}

func (self *LinkedList) FindLength() int {
	count := 0
	current := self.root
	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (self *LinkedList) PrintList() {
	if self.root == nil {
		fmt.Println("List is empty.")
		return
	}

	var node = self.root
	for node != nil {
		fmt.Printf("%s\n", node.data)
		node = node.next
	}
}
