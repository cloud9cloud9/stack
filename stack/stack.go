package st

import (
	"fmt"
	"strings"
)

type Stack[A any] struct {
	Head *Node[A]
}

type Node[A any] struct {
	Value A
	Next  *Node[A]
}

func (s *Stack[A]) Peek() (A, error) {
	if s.Head == nil {
		return s.Head.Value, fmt.Errorf("stack is empty")
	}
	return s.Head.Value, nil
}

func (s *Stack[A]) Pop() (A, error) {
	if s.Head == nil {
		return s.Head.Value, fmt.Errorf("stack is empty")
	}
	v := s.Head.Value
	s.Head = s.Head.Next
	return v, nil
}

func (s *Stack[A]) Push(el A) {
	node := &Node[A]{Value: el}
	node.Next = s.Head
	s.Head = node
}

func (s *Stack[A]) IsEmpty() bool {
	return s.Head == nil
}

func (s *Stack[A]) Size() int {
	size := 0
	for cur := s.Head; cur != nil; cur = cur.Next {
		size++
	}
	return size
}

func New[A any]() *Stack[A] {
	return &Stack[A]{}
}

func (s *Stack[A]) Print() {
	for s.Head != nil {
		fmt.Println(s.Head.Value)
		s.Head = s.Head.Next
	}
}

func (s Stack[A]) ToString() string {
	sb := new(strings.Builder)
	sb.WriteRune('{')

	for cur := s.Head; cur != nil; cur = cur.Next {
		sb.WriteString(fmt.Sprintf("%v,", cur.Value))
	}

	return strings.TrimRight(sb.String(), ",") + "}"
}
