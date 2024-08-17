package main

import (
	"fmt"
	st "stack/stack"
)

func main() {
	s := st.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())

	_, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Peek())
	fmt.Printf("Is stack empty: %t\n", s.IsEmpty())
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Print()
}
