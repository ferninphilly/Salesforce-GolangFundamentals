// my cool stack package
package stack

import "errors"

// A simple stack object with 10 slots for integers.
type Stack struct {
	ptr   int	// these fields are not exported
	items [10]int	// so they won't show up in the docs
}

// return true if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.ptr == 0
}

// return the number of items on the stack
func (s *Stack) Size() int {
	return s.ptr
}

// Push an item onto the stack
func (s *Stack) Push(i int) {
	s.items[s.ptr] = i
	s.ptr++
}

// Return topmost element, but do not pop it
func (s *Stack) Top() (int, error) {
        if s.IsEmpty() {
		return 0, errors.New("Stack empty!")
	}
	return s.items[s.ptr - 1], nil
}

// Pop an item from the stack
func (s *Stack) Pop() int {
	s.ptr--
	return s.items[s.ptr]
}
