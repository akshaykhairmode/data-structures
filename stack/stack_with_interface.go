package main

import "fmt"

type stacker interface {
	pop() (interface{}, error)
	push(interface{})
}

type stack struct {
	data []interface{}
}

type dummyStack struct {
	data []interface{}
}

func main() {

	s := &stack{}
	s.push(5)
	s.push(10)
	s.push(15)
	s.push(20)
	fmt.Println("Stack Before Pop", s)
	pop3Elements(s)
	fmt.Println("Stack After Pop", s)

	ds := &dummyStack{}
	ds.push(50)
	ds.push(100)
	ds.push("test")
	pop3Elements(ds)

}

func pop3Elements(s stacker) {
	for i := 0; i < 3; i++ {
		s.pop()
	}
}

func (this *stack) push(v interface{}) {
	this.data = append(this.data, v)
}

func (this *stack) pop() (interface{}, error) {

	if len(this.data) <= 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	elem := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return elem, nil
}

func (this *dummyStack) push(v interface{}) {
	fmt.Println("Dummy push method called")
}

func (this *dummyStack) pop() (interface{}, error) {

	fmt.Println("Dummy pop method called")

	return 0, nil
}
