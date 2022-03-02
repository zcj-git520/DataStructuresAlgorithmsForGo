package AVL_tree

import (
	"fmt"
	"testing"
)

func TestAVRTree_InsertData(t *testing.T) {
	avl := CreatAVLTree()
	avl.InsertData(2)
	avl.InsertData(12)
	avl.InsertData(24)
	avl.InsertData(1)
	avl.InsertData(65)
	avl.InsertData(72)
	avl.InsertData(32)
	avl.InsertData(21)
	avl.ShowNLR()
	avl.ShowLNR()
	avl.ShowLRN()
}

func TestAVLTree_Remove(t *testing.T) {
	avl := CreatAVLTree()
	avl.InsertData(1)
	avl.InsertData(2)
	avl.InsertData(3)
	avl.InsertData(4)
	avl.InsertData(5)
	avl.InsertData(6)
	avl.InsertData(7)
	avl.InsertData(8)
	avl.ShowNLR()
	avl.ShowLNR()
	avl.ShowLRN()
	fmt.Println("**********************")
	avl.Remove(4)
	avl.ShowNLR()
	avl.ShowLNR()
	avl.ShowLRN()
}

func TestAVLTree_IsValue(t *testing.T) {
	avl := CreatAVLTree()
	avl.InsertData(1)
	avl.InsertData(2)
	avl.InsertData(3)
	avl.InsertData(4)
	avl.InsertData(5)
	avl.InsertData(6)
	avl.InsertData(7)
	avl.InsertData(8)
	avl.ShowNLR()
	avl.ShowLNR()
	avl.ShowLRN()
	data := 4
	is := avl.IsValue(data)
	if is{
		fmt.Printf("%d :is exits\n", data)
	}else{
		fmt.Printf("%d :is not exits\n", data)
	}
}