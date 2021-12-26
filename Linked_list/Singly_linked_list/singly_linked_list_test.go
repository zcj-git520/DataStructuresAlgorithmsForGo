package Singly_linked_list

import (
	"fmt"
	"testing"
)

func TestLinkedList_AddToHead(t *testing.T) {
	newList := NewLinkerList()
	fmt.Println("list len is:", newList.Len())
	newList.AddToHead("stop")
	newList.AddToHead("2")
	newList.AddToHead("3")
	newList.AddToHead("4")
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
	newList.DeleteToHead()  // 头删除
	newList.DeleteToTail()  // 尾删除
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())

}

func TestLinkedList_AddToTail(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail("s1")
	//newList.AddToTail("2")
	//newList.AddToTail("3")
	//newList.AddToTail("3111")
	//newList.AddToTail(1)
	//newList.AddToTail(2)
	//newList.AddToTail("stop")
	newList.QuireAll()
	newList.DeleteToTail() // 尾删除
	newList.DeleteToTail() // 尾删除
	newList.DeleteToTail() // 尾删除
	newList.DeleteToTail() // 尾删除
	newList.DeleteToHead() // 头删除
	newList.QuireAll()

}

func TestLinkedList_AddListToHead(t *testing.T) {
	newList1 := NewLinkerList()
	newList2 := NewLinkerList()
	newList1.AddToTail(1)
	newList1.AddToTail(2)
	newList1.AddToTail(3)
	newList2.AddToTail(4)
	newList2.AddToTail(5)
	newList2.AddToTail(6)
	newList3 := newList2.AddListToHead(newList1)
	newList3.QuireAll()
	fmt.Println("list len is:", newList3.Len())
}

func TestLinkedList_AddListToTail(t *testing.T) {
	newList1 := NewLinkerList()
	newList2 := NewLinkerList()
	newList1.AddToTail(1)
	newList1.AddToTail(2)
	newList1.AddToTail(3)
	newList2.AddToTail(4)
	newList2.AddToTail(5)
	newList2.AddToTail(6)
	newList3 := newList1.AddListToTail(newList2)
	newList3.QuireAll()
	fmt.Println("list len is:", newList3.Len())
	newList4 := NewLinkerList()
	newList4.AddToTail(7)
	newList4.AddToTail(8)
	newList4.AddToTail(9)
	newList4.AddToTail(10)
	newList4 = newList3.AddListToHead(newList4)
	newList4.QuireAll()
	fmt.Println("list len is:", newList4.Len())
}

func TestLinkedList_AddListToIndex(t *testing.T) {
	newList1 := NewLinkerList()
	newList2 := NewLinkerList()
	newList1.AddToTail(1)
	newList1.AddToTail(2)
	newList1.AddToTail(3)
	newList2.AddToTail(4)
	newList2.AddToTail(5)
	newList2.AddToTail(6)
	newList3 := newList2.AddListToIndex(-3, newList1)
	newList3.QuireAll()
	fmt.Println("list len is:", newList3.Len())
}

func TestLinkedList_AddToIndex(t *testing.T) {
	newList := NewLinkerList()
	fmt.Println("list len is:", newList.Len())
	newList.AddToTail(0)
	newList.AddToTail(2)
	newList.AddToTail(3)
	newList.AddToTail(4)
	newList.AddToTail(5)
	newList.AddToIndex(1, 1)
	newList.AddToIndex(2, 10)
	newList.AddToIndex(-1, "stop")
	newList.AddToIndex(-2, "stop1")
	newList.AddToIndex(-3, "stop2")
	newList.AddToIndex(-6, "stop2")
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
}

func TestLinkedList_DeleteToAValue(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail(1)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(5)
	newList.AddToTail(6)
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
	fmt.Println("*******************************************")
	newList.DeleteToAValue(2)
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
}

func TestLinkedList_DeleteToValue(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail(1)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(2)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(3)
	newList.AddToTail(4)
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
	fmt.Println("*******************************************")
	newList.DeleteToValue(3)
	newList.QuireAll()
	fmt.Println("list len is:", newList.Len())
}

func TestLinkedList_DeleteToIndex(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail(0)
	newList.AddToTail(1)
	newList.AddToTail(2)
	newList.AddToTail(3)
	newList.AddToTail(4)
	newList.AddToTail(5)
	newList.AddToTail(6)
	newList.AddToTail(7)
	newList.AddToTail(8)
	newList.AddToTail(9)
	fmt.Println("list len is:", newList.Len())
	newList.QuireAll()
	newList.DeleteToIndex(0)
	fmt.Println("*********************************")
	fmt.Println("list len is:", newList.Len())
	newList.QuireAll()
}

func TestLinkedList_QuireValue(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail(1)
	newList.AddToTail(2)
	newList.AddToTail(3)
	newList.AddToTail(5)
	if !newList.QuireValue(2){
		t.Error("值不存在")
	}
}

func TestLinkedList_QuireIndex(t *testing.T) {
	newList := NewLinkerList()
	newList.AddToTail(0)
	newList.AddToTail(1)
	newList.AddToTail(2)
	newList.AddToTail(3)
	newList.AddToTail(4)
	newList.AddToTail(5)
	index := -1
	value := newList.QuireIndex(index)
	if value != 5{
		t.Error("index error", value)
	}
}
