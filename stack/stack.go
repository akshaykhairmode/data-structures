package main

import "fmt"

type stack struct {
	data []int
}

func main() {

	s := stack{}
	s.push(5)
	s.push(10)
	s.push(15)
	fmt.Println("Stack Before Pop", s)
	s.pop()
	s.pop()
	fmt.Println("Stack After Pop", s)

}

func (this *stack) push(v int) {
	this.data = append(this.data, v)
}

func (this *stack) pop() (int, error) {

	if len(this.data) <= 0 {
		return 0, fmt.Errorf("Stack is empty")
	}

	elem := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	return elem, nil
}
