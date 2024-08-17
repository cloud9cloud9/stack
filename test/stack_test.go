package test

import (
	"testing"

	st "stack/stack"

	"github.com/stretchr/testify/require"
)

func TestStackPush(t *testing.T) {
	s := st.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	require.Equal(t, 3, s.Size())
	require.Equal(t, "{3,2,1}", s.ToString())
}

func TestStackPop(t *testing.T) {
	s := st.New[int]()
	s.Push(1)
	s.Push(2)

	require.Equal(t, 2, s.Size())
	require.Equal(t, "{2,1}", s.ToString())
	_, err := s.Pop()
	_, err = s.Pop()

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	require.Equal(t, 0, s.Size())
	require.Equal(t, "{}", s.ToString())
	require.Equal(t, true, s.IsEmpty())
}

func TestStackPeek(t *testing.T) {
	s := st.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	require.Equal(t, "{3,2,1}", s.ToString())
	el, err := s.Peek()
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	require.Equal(t, 3, el)
}

func TestStackIsEmpty(t *testing.T) {
	s := st.New[int]()
	require.Equal(t, true, s.IsEmpty())
	s.Push(1)
	require.Equal(t, false, s.IsEmpty())
}
