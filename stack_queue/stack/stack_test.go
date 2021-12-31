package stack

import (
	"fmt"
	"testing"
)

func TestNewStackData(t *testing.T) {
	s := NewStackData()
	s.Push(0)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.ShowStack()
	fmt.Println("*************************************")
	data := s.GetTopValue()
	fmt.Println("data = >", data)
	data = s.GetTopValue()
	fmt.Println("data = >", data)
	data = s.GetTopValue()
	fmt.Println("data = >", data)
	s.ShowStack()
	fmt.Println("************************************")
	data = s.Pop()
	fmt.Println("data = >", data)
	s.ShowStack()
	fmt.Println("************************************")
	data = s.Pop()
	fmt.Println("data = >", data)
	s.ShowStack()
	fmt.Println("************************************")
	data = s.Pop()
	fmt.Println("data = >", data)
	s.ShowStack()
	fmt.Println("************************************")
	fmt.Println("stack len is ", s.Len())

}
