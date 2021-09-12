package main

import "fmt"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	next *Node
	prev *Node
	data interface{}
}

func (this *DoublyLinkedList) Prepend(n *Node) {

	if n == nil {
		return
	}

	defer func() {
		this.length++
	}()

	//If first node then its both head and tail
	if this.length <= 0 {
		this.tail = n
		this.head = n
		return
	}

	//Since old head will become the next node, we will set the previous as passed node
	oldHead := this.head
	oldHead.prev = n

	//We now say n is the new head and new head will point to old head as next
	this.head = n
	this.head.next = oldHead

}

func (this *DoublyLinkedList) Append(n *Node) {

	if n == nil {
		return
	}

	defer func() {
		this.length++
	}()

	//If first node then its both head and tail
	if this.length <= 0 {
		this.head = n
		this.tail = n
		return
	}

	//Previous node will now point to passed node
	previousNode := this.tail
	previousNode.next = n

	//Set new tail as the passed node
	this.tail = n
	n.prev = previousNode
}

func (this *DoublyLinkedList) Remove(n *Node) {

	if n == nil {
		return
	}

	defer func() {
		this.length--
	}()

	//If we are removing the head node then we say that the next node from head is the head node now
	if n == this.head {
		this.head = this.head.next
		this.head.prev = nil
		return
	}

	//If we are removing the tail node then we say that the previous node from tail node is the tail node now
	if n == this.tail {
		this.tail = this.tail.prev
		this.tail.next = nil
		return
	}

	//If its a middle node then, previous nodes next will point to next node and next nodes previous will point to previous node
	prev := n.prev
	next := n.next

	prev.next = next
	next.prev = prev

}

func (this *DoublyLinkedList) ShiftRightBy1() {

	previousNode := this.tail.prev

	this.tail.prev = nil
	this.tail.next = this.head

	//Modify Head
	this.head.prev = this.tail

	//Finally Set the tail as head, this will move tail to head
	this.head = this.tail
	this.tail = previousNode
	this.tail.next = nil
}

func (this *DoublyLinkedList) ShiftLeftBy1() {

	//Modify the next node, the node which will become head
	nextNode := this.head.next
	nextNode.prev = nil

	//Get the lastnode and set head as its next
	this.tail.next = this.head

	//Since we will move head to the previous
	this.head.prev = this.tail
	this.head.next = nil

	//New head will be the next node
	this.head = nextNode
	this.tail = this.head
}

func (this DoublyLinkedList) PrintList() {
	curr := this.head
	for this.length != 0 {
		fmt.Print(curr.data, " ")
		curr = curr.next
		this.length--
	}
	fmt.Println()
}

func (this DoublyLinkedList) PrintAdjacentNodes() {
	curr := this.head
	for this.length != 0 {
		fmt.Print(curr.data)
		if curr.prev != nil {
			fmt.Print(" - Prev : ", curr.prev.data)
		} else {
			fmt.Print(" - Prev : ", nil)
		}

		if curr.next != nil {
			fmt.Print(" | Next : ", curr.next.data)
		} else {
			fmt.Print(" | Next : ", nil)
		}

		curr = curr.next
		this.length--
		fmt.Println()
	}
}

func main() {

	dll := DoublyLinkedList{}
	n1 := &Node{data: "first"}
	n2 := &Node{data: "second"}
	n3 := &Node{data: "third"}
	n4 := &Node{data: "fourth"}
	n5 := &Node{data: "fifth"}
	dll.Prepend(n1)
	dll.Prepend(n2)
	dll.Prepend(n3)
	dll.Append(n4)
	dll.Prepend(n5)
	dll.Remove(n2)
	dll.PrintAdjacentNodes()
	dll.ShiftRightBy1()
	dll.ShiftLeftBy1()
	fmt.Println("============================")
	dll.PrintAdjacentNodes()

}
