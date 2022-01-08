package binary_search_tree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree_InsertNode(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.InsertNode(7)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(12)
	tree.InsertNode(8)
	tree.InsertNode(4)
	tree.ShowNLR()
	tree.ShowLNR()
	tree.ShowLRN()
}

func TestBinarySearchTree_RemoveByMerge(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.InsertNode(7)
	tree.InsertNode(3)
	tree.InsertNode(5)
	tree.InsertNode(-12)
	tree.InsertNode(15)
	tree.InsertNode(12)
	tree.InsertNode(8)
	tree.InsertNode(4)
	tree.ShowNLR()
	tree.ShowLNR()
	tree.ShowLRN()
	fmt.Println("*******************************************")
	fmt.Println(tree.RemoveByMerge(15))
	tree.ShowNLR()
	tree.ShowLNR()
	tree.ShowLRN()
	fmt.Println("*******************************************")
}

func TestBinarySearchTree_RemoveByCopy(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.InsertNode(7)
	tree.InsertNode(3)
	tree.InsertNode(5)
	//tree.InsertNode(-12)
	//tree.InsertNode(15)
	tree.InsertNode(12)
	tree.InsertNode(8)
	tree.InsertNode(4)
	tree.ShowNLR()
	tree.ShowLNR()
	tree.ShowLRN()
	fmt.Println(tree.GetNodeNum())
	fmt.Println("*******************************************")
	fmt.Println(tree.RemoveByCopy(8))
	tree.ShowNLR()
	tree.ShowLNR()
	tree.ShowLRN()
	fmt.Println(tree.GetNodeNum())
	fmt.Println("*******************************************")
}

func TestBinarySearchTree_IsValue(t *testing.T) {
	tree := NewBinarySearchTree()
	tree.InsertNode(7)
	tree.InsertNode(3)
	tree.InsertNode(5)
	//tree.InsertNode(-12)
	//tree.InsertNode(15)
	tree.InsertNode(12)
	tree.InsertNode(8)
	tree.InsertNode(4)
	data := 76
	if tree.IsValue(data){
		fmt.Println(data , "is exits")
	}else{
		fmt.Println(data , "is not  exits")
	}
}