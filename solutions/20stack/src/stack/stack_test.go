package stack

//import "stack" // no need to import if in same dir
import "testing"
import "fmt"

func TestPushPop(t *testing.T) {
	s := Stack{} // s := stack.Stack{}
	s.Push(12)
	fmt.Printf("%#v\n", s)
	if s.Pop() != 12 {
		t.Log("Pop() failed to return 1")
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	s := Stack{} // s := stack.Stack{}
	if s.Size() != 0 {
		t.Log("Size() failed to return 0 for empty stack!")
		t.Fail()
	}
}

func TestIsEmpty(t *testing.T) {
	s := Stack{} // s := stack.Stack{}
	if !s.IsEmpty() {
		t.Log("IsEmpty() failed to return true for empty stack!")
		t.Fail()
	}
}

func TestTop(t *testing.T) {
	s := Stack{} // s := stack.Stack{}
	s.Push(-2)
	ret, _ := s.Top()

	if ret != -2 {
		t.Log("Top() failed to return item on top of stack!")
		t.Fail()
	}
}
func ExamplePush() {
	s := Stack{} // s := stack.Stack{}
	s.Push(1)
	fmt.Println(s.Pop())
	// Output:
	// 1
}
