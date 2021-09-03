package main

import "fmt"

type LinkedList struct {
	head   *Node
	length int
}

type Node struct {
	next *Node
	data interface{}
}

func (this *LinkedList) Prepend(n *Node) {
	oldHead := this.head
	this.head = n
	this.head.next = oldHead
	this.length++
}

func (this *LinkedList) Remove(n *Node) {

	l := this.length
	var curr, prev *Node

	defer func() {
		this.length--
	}()

	if this.head == n {
		this.head = this.head.next
		return
	}

	curr = this.head.next
	prev = this.head

	for l != 0 {
		if curr == n {
			prev.next = curr.next
			return
		} else {
			prev = curr
			curr = curr.next
		}
		l--
	}

}

func (this LinkedList) PrintList() {
	curr := this.head
	for this.length != 0 {
		fmt.Println(curr.data)
		curr = curr.next
		this.length--
	}
}

func main() {

	ll := LinkedList{}
	n1 := &Node{data: "first"}
	n2 := &Node{data: "second"}
	n3 := &Node{data: "third"}
	ll.Prepend(n1)
	ll.Prepend(n2)
	ll.Prepend(n3)
	ll.Remove(n1)
	ll.Remove(n2)
	ll.PrintList()

}
