package stacks

import "fmt"

type stacks struct {
	index int
	stack []any
}

func NewStack() *stacks {
	return &stacks{
		index: -1,
	}
}

func (s *stacks) Push(value any) {
	s.stack = append(s.stack, value)
	s.index++
}

func (s *stacks) Pop() (value any, err error) {
	if s.index < 0 {
		err = fmt.Errorf("Stack is empty")
		return
	}

	value = s.stack[s.index]
	s.index--
	return
}

func (s *stacks) Peek() (value any, err error) {
	if s.index < 0 {
		err = fmt.Errorf("Stack is empty")
		return
	}

	value = s.stack[s.index]
	return
}

func (s *stacks) Size() int {
	return s.index + 1
}
