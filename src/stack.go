package stack

// A simple stack object with 10 slots for integers
type Stack struct {
    ptr  int
    items [10]int
}
// Push an item onto the stack
func (s *Stack) Push(i int) {
    s.items[s.ptr] = i
    s.ptr++
}
// Pop an item from the stack
func (s *Stack) Pop() int {
    s.ptr--
    return s.items[s.ptr]
}
