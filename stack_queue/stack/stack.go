/*
author：zhaochengji mail：909536346@qq.com
stack for go
栈，基于单向链表
*/
package stack

import "github.com/zcj-git520/DataStructuresAlgorithmsForGo/Linked_list/Singly_linked_list"

type stackData struct {
	list *Singly_linked_list.LinkedList
}

// 获得栈的长度
func (s *stackData)Len() int{
	return s.list.Len()
}

// 将数据插入栈顶
func (s *stackData)Push(data interface{}) {
	s.list.AddToHead(data)
}

// 将数据从栈顶取出，并删除数据
func (s *stackData)Pop()interface{}{
	data := s.list.QuireIndex(0)
	// 栈不为空, 删除栈顶数据
	if data != nil{
		s.list.DeleteToHead()
	}
	return data
}

// 将数据取出，不删除数据
func (s *stackData)GetTopValue()interface{}{
	return s.list.QuireIndex(0)
}

// 展示栈
func (s *stackData)ShowStack() {
	s.list.QuireAll()
}

// 创建栈
func NewStackData()*stackData{
	return  &stackData{list:Singly_linked_list.NewLinkedList()}
}
