/*
author：zhaochengji mail：909536346@qq.com
circle_double_linked_listt for go
链表：是由若干的节点组成的链式结构，每个节点包含着存储的信息和指向下一节点的指针
循环双向链表
*/
package circle_double_linked_list

import "fmt"

// 节点的定义
type circleDoubleNode struct {
	info  interface{}  // 数据结构的定义
	prev  *circleDoubleNode  // 前驱指针
	next  *circleDoubleNode  // 后驱指针
}

// 双向循环链表的定义
type circleDoubleList struct {
	currentNode *circleDoubleNode  // 指向链表的指针
	len  int    // 链表的长度
}

// 创建节点
func createNode(data interface{}) *circleDoubleNode{
	return &circleDoubleNode{
		info: data,
		prev: nil,
		next: nil,
	}
}

// 得到链表的长度
func (c *circleDoubleList)GetLen()int{
	return c.len
}

// 节点的插入
func (c *circleDoubleList)InsertNode(data interface{}) {
	newNode := createNode(data)
	// 如果链表为空
	if c.currentNode == nil{
		c.currentNode = newNode  // 新创建的节点为当前节点
		// 节点的前驱与后驱都指向自己
		c.currentNode.next = c.currentNode
		c.currentNode.prev = c.currentNode
	}else{
		// 节点的插入
		c.currentNode.next.prev = newNode
		newNode.prev = c.currentNode
		newNode.next = c.currentNode.next
		c.currentNode.next = newNode
		c.currentNode = newNode
	}
	// 链表长度 +1
	c.len ++
}

// 当前节点的删除
func (c *circleDoubleList)DeleteNode() bool{
	// 链表为空
	if c.currentNode == nil{
		return false
	}
	// 当链表值存在一个元素时
	if c.len == 1{
		c.currentNode = nil
		c.len --
		return true
	}
	// 当链表只存在两个元素时
	if c.len == 2{
		c.currentNode = c.currentNode.next
		c.len --
		return true
	}
	// 当前节点的后驱节点的前驱指针指向当前节点的的前驱节点
	c.currentNode.next.prev = c.currentNode.prev
	// 当前节点的前驱节点的后驱指针指向当前节点的后驱节点
	c.currentNode.prev.next = c.currentNode.next
	// 设置当前节点的后驱节点为当前节点
	c.currentNode = c.currentNode.next
	c.len --
	return true
}

// 遍历说有节点
func (c *circleDoubleList)QuireAll(){
	if c.currentNode == nil{
		fmt.Println("circle double linked list is empty")
		return
	}
	//__index := 1
	//for node := c.currentNode.next; node != c.currentNode; node = node.next{
	//	fmt.Printf("index is %d, node value is %v", __index, node.info)
	//	__index ++
	//}
	fmt.Printf("circle double linked list len is %d \n", c.len)
	node := c.currentNode.next
	for i := 1; i < c.len; i++{
		fmt.Printf("index is %d, node value is %v \n", i, node.info)
		node = node.next
	}
	fmt.Printf("index is %d, node value is %v \n", c.len, c.currentNode.info)
}

// 判断值是否存在
func (c *circleDoubleList)QuireValue(data interface{}) bool{
	if c.currentNode == nil{
		return false
	}
	node := c.currentNode
	for i := 1; i <= c.len; i++{
		if node.info == data{
			return true
		}
		node = node.next
	}
	return  false
}

// 链表实例化
func NewCircleDoubleList()*circleDoubleList{
	return &circleDoubleList{
		currentNode: nil,
		len:         0,
	}
}