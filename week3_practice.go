package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type StackMethods interface {
	Peek() (int, error)
	Pop() error
	Push(int)
	IsEmpty() bool
	Print()
}

type StackIncrement interface {
	Increment(int)
}

type Stack struct {
	Top *Node
}

func (s *Stack) Push(v int) (int, error) {
	var StartValue int
	temp := Node{v, nil}
	if s.Top == nil {
		s.Top = &temp
		return 0, fmt.Errorf("No element before")
	}
	StartValue = s.Top.Value
	temp.Next = s.Top
	s.Top = &temp

	return StartValue, nil
}

func (s *Stack) Pop() error {
	if s.Top == nil {
		return fmt.Errorf("No elements in Stack")
	}
	s.Top = s.Top.Next
	return nil
}

func (s *Stack) IsEmpty() bool {
	if s.Top == nil {
		return true
	}
	return false
}

func (s *Stack) Print() {
	temp := s.Top
	for temp != nil {
		fmt.Println(temp.Value)
		temp = temp.Next
	}
}

func (s Stack) Peek() (int, error) {
	if s.Top == nil {
		return 0, fmt.Errorf("Empty Stack")
	}
	return s.Top.Value, nil
}

func (s Stack) Increment(v int) {
	temp := s.Top
	for v != 0 && temp != nil {
		temp.Value += 1
		temp = temp.Next
		v--
	}
}

func main() {
	s := Stack{}
	fmt.Println(s)
	s.Push(1)
	fmt.Println(s.Top.Value)

	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Print()

	fmt.Println("")
	s.Pop()
	s.Print()

	fmt.Println(s.IsEmpty())

	s.Increment(5)
	s.Print()

	fmt.Println("")
	fmt.Println(s.Peek())

	s.Pop()
	s.Pop()
	s.Pop()
	fmt.Println("")
	fmt.Println(s.Peek())
}
