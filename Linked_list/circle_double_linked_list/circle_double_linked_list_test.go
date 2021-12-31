package circle_double_linked_list

import (
	"fmt"
	"testing"
)

func TestCircleDoubleList_InsertNode(t *testing.T) {
	newList := NewCircleDoubleList()
	newList.InsertNode(0)
	newList.InsertNode(1)
	newList.InsertNode(2)
	newList.InsertNode(3)
	newList.InsertNode(4)
	newList.QuireAll()
}

func TestCircleDoubleList_DeleteNode(t *testing.T) {
	newList := NewCircleDoubleList()
	newList.InsertNode(0)
	newList.InsertNode(1)
	newList.InsertNode(2)
	newList.InsertNode(3)
	newList.InsertNode(4)
	newList.QuireAll()
	fmt.Println("************************************************")
	newList.DeleteNode()
	newList.DeleteNode()
	newList.DeleteNode()
	newList.QuireAll()
}

func TestCircleDoubleList_QuireValue(t *testing.T) {
	newList := NewCircleDoubleList()
	newList.InsertNode(0)
	newList.InsertNode(1)
	newList.InsertNode(2)
	newList.InsertNode(3)
	newList.InsertNode(4)
	newList.QuireAll()
	fmt.Println("*********************************************")
	data := 1
	if newList.QuireValue(data){
		fmt.Printf("%d is exists\n", data)
	}else{
		fmt.Printf("%d is not exists\n", data)
	}
}