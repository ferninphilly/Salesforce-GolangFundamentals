package stack

import "testing"

func TestPushPop(t *testing.T) {
        s := Stack{} // cf. s := new(Stack)
        s.Push(1)
        if s.Pop() != 1 {
                t.Log("Pop() failed to return 1")
                t.Fail()
        }
}
