/*
author：zhaochengji mail：909536346@qq.com
singly_linked_list for go
链表：是由若干的节点组成的链式结构，每个节点包含着存储的信息和指向下一节点的指针
单向链表：节点中只有指向后继节点的指针
*/
package Singly_linked_list

import (
	"fmt"
	"math"
)

// 单项链表节点的定义
type singlyLinkedNode struct {
	info    interface{}    		// 存储任意数据
	next    *singlyLinkedNode  	// 指向下一节点
}

type LinkedList struct {
	len  int                   // 链表长度
	Head *singlyLinkedNode     // 头节点
	Tail *singlyLinkedNode     // 尾节点
}

// 创建节点
func createNode(info interface{}) *singlyLinkedNode {
	node := &singlyLinkedNode{
		info: info,
		next: nil,
	}
	return node
}

// 链表的长度
func(l *LinkedList)Len() int{
	return l.len
}

/*
单个节点的插入
*/
// 链表的尾插法
func (l *LinkedList)AddToTail(info interface{}) *LinkedList {
	newNode := createNode(info)  // 创建新节点
	// 链表为空, 头尾节点都指向该节点
	if l.Tail == nil{
		l.Head = newNode
		l.Tail = newNode
	}else {
		l.Tail.next = newNode
		l.Tail = l.Tail.next
	}
	// 链表长度+1
	l.len++
	return l
}

// 链表的头插法
func (l *LinkedList)AddToHead(info interface{}) *LinkedList  {
	newNode := createNode(info)  // 创建节点
	// 链表为空
	if l.Head == nil{
		l.Head = newNode
		l.Tail = newNode
	}else {
		newNode.next = l.Head
		l.Head = newNode
	}
	// 链表长度+1
	l.len++
	return  l
}

// 链表的index
// 通过index=>插入链表新的节点
// index为正数 为从左->右 // 0表示第一个节点
// index为负数 为从右->左 // -1表示第一个节点
func (l *LinkedList)AddToIndex(index int, info interface{}) *LinkedList {
	// 当链表为空时，采用了尾插入法插入数据
	if l.Head == nil{
		return l.AddToTail(info)
	}
	// 当index大于链表的值时，默认将数据插入到链表的后面
	if int(math.Abs(float64(index))) > l.len-1{
		return l.AddToTail(info)
	}
	// 插入链表的头部
	if index == 0{
		return l.AddToHead(info)
	}
	// 插入到链表的末尾
	if index == -1{
		return l.AddToTail(info)
	}
	// 当index 为负数时，从右到左插入
	if index < 0{
		index += l.len
	}
	__index := 1  // 内部的index值
	// 从第二个值开始插入
	for node := l.Head.next; node != l.Tail; node = node.next{
		if index == __index{
			newNode := createNode(info)
			newNode.next = node.next    // 指向node的后继节点
			node.next = newNode         // node 指向这个新的节点
			// 链表的节点数加1
			l.len ++
			return l
		}
		__index ++
	}
	return l
}

/*
链表的合并
*/
// 将新的链表插入头部
func (l *LinkedList)AddListToHead(list *LinkedList)*LinkedList{
	// 两个链表都为空时
	if l.Head == nil && list.Head == nil{
		return l
	}
	// 合并表为空
	if l.Head != nil && list.Head == nil{
		return l
	}
	// 主表为空
	if l.Head == nil && list.Head != nil{
		return list
	}
	// 合并表的尾指针指向主表的头
	list.Tail.next = l.Head
	list.Tail = l.Tail
	// 表节点数合并
	list.len += l.len
	return list
}

// 将新的链表插入到尾部
func (l *LinkedList)AddListToTail(list *LinkedList)*LinkedList{
	// 两个链表都为空时
	if l.Head == nil && list.Head == nil{
		return l
	}
	// 合并表为空
	if l.Head != nil && list.Head == nil{
		return l
	}
	// 主表为空
	if l.Head == nil && list.Head != nil{
		return list
	}
	// 将新表插入到主表之后
	l.Tail.next = list.Head
	l.Tail = list.Tail
	l.len += list.len
	return l
}

// 经新的表插入到index
func (l *LinkedList)AddListToIndex(index int, list *LinkedList) *LinkedList {
	// 两个链表都为空时
	if l.Head == nil && list.Head == nil{
		return l
	}
	// 合并表为空
	if l.Head != nil && list.Head == nil{
		return l
	}
	// 主表为空
	if l.Head == nil && list.Head != nil{
		return list
	}
	if int(math.Abs(float64(index))) > l.len{
		fmt.Println("错误的index")
		return l
	}
	if index == 0{
		return l.AddListToHead(list)
	}
	if int(math.Abs(float64(index))) == l.len -1 || index == -1{
		return l.AddListToTail(list)
	}
	if index < 0{
		index += l.len
	}
	__index := 1
	for node := l.Head.next; node != l.Tail; node = node.next{
		if index == __index{
			list.Tail.next = node.next
			node.next = list.Head
			l.len += list.len  // 链表的数值相加
			return l
		}
		__index ++
	}
	return l
}

// 链表头删除法
func (l *LinkedList)DeleteToHead()*LinkedList{
	// 链表为空
	if l.Head == nil{
		return l
	}
	// 链表中只有一个数
	if l.Head == l.Tail{
		l.Head = nil
		l.Tail = nil
		l.len  = 0
		return l
	}
	l.Head = l.Head.next
	// node数减一
	l.len --
	return l
}

// 链表尾删除法
func (l *LinkedList)DeleteToTail()*LinkedList{
	// 链表为空
	if l.Tail == nil{
		return l
	}
	// 链表中只有一个数
	if l.Head == l.Tail{
		l.Head = nil
		l.Tail = nil
		l.len = 0
		return l
	}
	// 找出尾节点的上一节点
	var node  *singlyLinkedNode
	for node = l.Head;  node.next != l.Tail; node = node.next{}
	node.next = nil
	l.Tail = node
	// node数减一
	l.len --
	return l
}

// 通过值=>删除链表的节点(第一个)
func (l *LinkedList)DeleteToAValue(value interface{})*LinkedList{
	// 链表为空
	if l.Head == nil{
		fmt.Println("链表为空")
		return l
	}
	// value == head.info
	// 就采用头删法
	if value == l.Head.info{
		return  l.DeleteToHead()
	}
	// 中间采用轮询查找value, 从第二个开始轮询到倒数第二个结束
	node := l.Head
	for temp := l.Head.next; temp != nil; temp = node.next{
		if temp.info == value{
			//  删除node
			node.next = temp.next
			l.len --
			return l
		}
		node = node.next
	}
	fmt.Println("链表中：值不存在")
	return l
}

// 通过值=>删除链表的节点(所有)
func (l *LinkedList)DeleteToValue(value interface{})*LinkedList{
	// 链表为空
	if l.Head == nil{
		fmt.Println("链表为空")
		return l
	}
	node := l.Head         // 当前的node
	temp := l.Head.next    // 下一个node
	for {
		// 当下一个节点是尾节点时，就判断首位和末尾是为需要删除的node
		if temp == l.Tail {
			if l.Head.info == value{
				l.DeleteToHead()
			}
			if l.Tail.info == value{
				l.DeleteToTail()
			}
			return l
		}
		if temp.info == value{
			// 删除node
			node.next = temp.next
			temp = temp.next
			// node数减一
			l.len --
			continue
		}
		node = node.next
		temp = temp.next
		}
}

// 通过index=>删除链表的节点
// index为正数 为从左->右 // 0表示第一个节点
// index为负数 为从右->左 // -1表示第一个节点
func (l *LinkedList)DeleteToIndex(index int)*LinkedList{
	// 链表为空
	if l.Head == nil{
		fmt.Println("链表为空")
		return l
	}
	// index 超过链表数
	if int(math.Abs(float64(index))) > l.len-1{
		fmt.Println("错误的index")
		return l
	}
	// 删除第一个node
	if index == 0{
		return l.DeleteToHead()
	}
	if index < 0{
		index += l.len
	}
	_index := 1
	node := l.Head
	temp := node.next
	for {
		if index == _index{
			if temp != nil{
				node.next = temp.next
			}
			l.len --
			return l
		}
		node = node.next
		temp = temp.next
		_index ++
	}
}

// 遍历链表
func (l *LinkedList) QuireAll() {
	if l.Head == nil{
		fmt.Println("链表数据为空")
		return
	}
	node := l.Head
	for {
		if node == nil{
			return
		}
		fmt.Println(node.info)
		if node == l.Tail{
			return
		}
		node = node.next
	}
}

// 判断valve是否存在
func (l *LinkedList)QuireValue(value interface{}) bool{
	// 表为空
	if l.Head == nil{
		return false
	}
	// 表中只存在一个值
	if l.Head == l.Tail{
		if l.Head.info == value{
			return true
		}
		return false
	}
	// 遍历查值
	for node := l.Head; node != l.Tail.next; node = node.next{
		if node.info == value{
			return true
		}
	}
	return false
}

// 根据索引返回值
func (l *LinkedList)QuireIndex(index int) interface{} {
	// 链表为空
	if l.Head == nil{
		return nil
	}
	if int(math.Abs(float64(index))) > l.len -1 {
		fmt.Println("索引值错误")
		return nil
	}
	if index == 0{
		return l.Head.info
	}
	if index == l.len -1 || index == -1{
		return l.Tail.info
	}
	if index < 0{
		index += l.len
	}
	__index := 1
	for node := l.Head.next; node != l.Tail; node = node.next{
		if __index == index{
			return node.info
		}
		__index ++
	}
	//fmt.Println("未能找到！！！！")
	return nil
}

func NewLinkedList()*LinkedList{
	// 创建的链表头尾节点都为空
	return &LinkedList{
		Head: nil,
		Tail: nil,
		len: 0,
	}
}

