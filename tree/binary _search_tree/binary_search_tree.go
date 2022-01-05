/*
author：zhaochengji mail：909536346@qq.com
binary search tree for go
*/
package binary__search_tree


// 定义树的节点
type binarySearchTreeNode struct {
	info  interface{}             // 存储的数据
	left  *binarySearchTreeNode   // 左节点
	right *binarySearchTreeNode   // 右节点
}

// 定义树结构
type  BinarySearchTree struct {
	root  *binarySearchTreeNode  // 定义头节点
	treeNodeNum  int             // 节点的数量
}

// 初始节点
func newTreeNode(info interface{}) *binarySearchTreeNode {
	return &binarySearchTreeNode{
		info:  info,
		left:  nil,
		right: nil,
	}
}

// 插入数据
func (b *BinarySearchTree)insertNode(data interface{})  {
	// 树为空树
	newNode := newTreeNode(data)
	if b.root == nil{
		b.root = newNode
		b.treeNodeNum++
		return
	}
	node := b.root
	for{
		if node == nil{
			return
		}
		if node.info <= data{
			if node.right == nil{
				node.right = newNode
				b.treeNodeNum ++
				return
			}else {
				node = node.right
			}
		} else{
			if node.left == nil{
				node.left = node
				b.treeNodeNum ++
				return
			}else {
				node = node.left
			}

		}
	}
}