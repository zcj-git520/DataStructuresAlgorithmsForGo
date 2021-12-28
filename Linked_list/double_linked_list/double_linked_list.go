/*
author：zhaochengji mail：909536346@qq.com
singly_linked_list for go
链表：是由若干的节点组成的链式结构，每个节点包含着存储的信息和指向下一节点的指针
双向链表： 节点中存在两个指针，一个指针指向先驱节点，另一个指向后继节点
*/
package double_linked_list

import (
	"fmt"
	"math"
)

// 双向链表的节点定义
type doubleLinkedNode struct {
	info  interface{}       // 存储的数据
	prev *doubleLinkedNode  // 指向先驱节点
	next *doubleLinkedNode  // 指向后驱几点
}

// 双向链表定义
type doubleList struct {
	Head *doubleLinkedNode   // 头节点
	Tail *doubleLinkedNode   // 尾节点
	len  int                 // 链表长度
}

// 创建节点
func createNode(data interface{}) *doubleLinkedNode {
	return &doubleLinkedNode{
		info: data,
		prev: nil,
		next: nil,
	}
}

// 链表的长度
func(d *doubleList)Len() int{
	return d.len
}

/*
单个节点的插入
*/
// 链表的尾插法
func (d *doubleList)AddToTail(info interface{}) *doubleList {
	newNode := createNode(info)  // 创建新节点
	// 链表为空, 头尾节点都指向该节点
	if d.Head == nil{
		d.Head = newNode
		d.Tail = newNode
	}else{
		d.Tail.next = newNode
		newNode.prev = d.Tail
		d.Tail = d.Tail.next
	}
	// 链表长度+1
	d.len++
	return d
}

// 链表的头插法
func (d *doubleList)AddToHead(info interface{}) *doubleList  {
	newNode := createNode(info)  // 创建节点
	// 链表为空
	if d.Head == nil{
		d.Head = newNode
		d.Tail = newNode
	}else{
		newNode.next = d.Head    // 新节点的后驱指向头节点
		d.Head.prev = newNode    // 头节点的前驱指向新节点
		d.Head = newNode         // 将新节点设置尾头节点
	}
	// 链表长度+1
	d.len++
	return  d
}

// 链表的index
// 通过index=>插入链表新的节点
// index为正数 为从左->右 // 0表示第一个节点
// index为负数 为从右->左 // -1表示第一个节点
func (d *doubleList)AddToIndex(index int, info interface{}) *doubleList {
	// 当链表为空时，采用了尾插入法插入数据
	if d.Head == nil{
		return d.AddToTail(info)
	}
	// 当index大于链表的值时，默认将数据插入到链表的后面
	if int(math.Abs(float64(index))) > d.len-1{
		return d.AddToTail(info)
	}
	// 插入链表的头部
	if index == 0{
		return d.AddToHead(info)
	}
	// 插入到链表的末尾
	if index == -1{
		return d.AddToTail(info)
	}
	// 当index 为负数时，从右到左插入
	if index < 0{
		index += d.len +1
	}
	__index := 1  // 内部的index值
	// 从第二个值开始插入
	for node := d.Head; node != d.Tail; node = node.next{
		if index == __index{
			newNode := createNode(info)
			node.next.prev = newNode   // 节点的后驱节点的先驱指向新节点
			newNode.prev = node        // 新节点的前驱指向节点
			newNode.next = node.next   // 新节点的后驱指向节点的后驱
			node.next = newNode           // 节点的后驱指向新节点
			// 链表的节点数加1
			d.len ++
			return d
		}
		__index ++
	}
	return d
}

/*
链表的合并
*/
// 将新的链表插入头部
func (d *doubleList)AddListToHead(list *doubleList)*doubleList{
	// 两个链表都为空时
	if d.Head == nil && list.Head == nil{
		return d
	}
	// 合并表为空
	if d.Head != nil && list.Head == nil{
		return d
	}
	// 主表为空
	if d.Head == nil && list.Head != nil{
		return list
	}
	// 合并表的尾指针指向主表的头
	list.Tail.next = d.Head
	d.Head.prev = list.Tail
	list.Tail = d.Tail
	// 表节点数合并
	list.len += d.len
	return list
}

// 将新的链表插入到尾部
func (d *doubleList)AddListToTail(list *doubleList)*doubleList{
	// 两个链表都为空时
	if d.Head == nil && list.Head == nil{
		return d
	}
	// 合并表为空
	if d.Head != nil && list.Head == nil{
		return d
	}
	// 主表为空
	if d.Head == nil && list.Head != nil{
		return list
	}
	// 将新表插入到主表之后
	d.Tail.next = list.Head
	list.Head.prev = d.Tail
	d.Tail = list.Tail
	d.len += list.len
	return d
}

// 经新的表插入到index
func (d *doubleList)AddListToIndex(index int, list *doubleList) *doubleList {
	// 两个链表都为空时
	if d.Head == nil && list.Head == nil{
		return d
	}
	// 合并表为空
	if d.Head != nil && list.Head == nil{
		return d
	}
	// 主表为空
	if d.Head == nil && list.Head != nil{
		return list
	}
	if int(math.Abs(float64(index))) > d.len{
		fmt.Println("错误的index")
		return d
	}
	if index == 0{
		return d.AddListToHead(list)
	}
	if  index == -1{
		return d.AddListToTail(list)
	}
	if index < 0{
		index += d.len
	}
	__index := 1
	for node := d.Head; node != d.Tail; node = node.next{
		if index == __index{
			node.next.prev = list.Tail
			list.Tail.next = node.next
			node.next = list.Head
			list.Head.prev = node
			d.len += list.len  // 链表的数值相加
			return d
		}
		__index ++
	}
	return d
}

// 链表头删除法
func (d *doubleList)DeleteToHead()*doubleList{
	// 链表为空
	if d.Head == nil{
		return d
	}
	// 链表中只有一个数
	if d.Head == d.Tail{
		d.Head = nil
		d.Tail = nil
		d.len  = 0
		return d
	}
	d.Head = d.Head.next
	d.Head.prev = nil
	// node数减一
	d.len --
	return d
}

// 链表尾删除法
func (d *doubleList)DeleteToTail()*doubleList{
	// 链表为空
	if d.Tail == nil{
		return d
	}
	// 链表中只有一个数
	if d.Head == d.Tail{
		d.Head = nil
		d.Tail = nil
		d.len = 0
		return d
	}
	d.Tail = d.Tail.prev
	d.Tail.next = nil
	// node数减一
	d.len --
	return d
}

// 通过值=>删除链表的节点(第一个)
func (d *doubleList)DeleteToAValue(value interface{})*doubleList{
	// 链表为空
	if d.Head == nil{
		fmt.Println("链表为空")
		return d
	}
	// value == head.info
	// 就采用头删法
	if value == d.Head.info{
		return  d.DeleteToHead()
	}
	if value == d.Tail.info{
		return  d.DeleteToTail()
	}
	// 中间采用轮询查找value, 从第二个开始轮询到倒数第二个结束
	for node := d.Head.next; node != d.Tail; node = node.next{
		if node.info == value{
			//  删除node
			node.next.prev = node.prev
			node.prev.next = node.next
			d.len --
			return d
		}
	}
	fmt.Println("链表中：值不存在")
	return d
}

// 通过值=>删除链表的节点(所有)
func (d *doubleList)DeleteToValue(value interface{})*doubleList{
	// 链表为空
	if d.Head == nil{
		fmt.Println("链表为空")
		return d
	}
	node := d.Head         // 当前的node
	for {
		// 当下一个节点是尾节点时，就判断首位和末尾是为需要删除的node
		if node == d.Tail {
			if d.Head.info == value{
				d.DeleteToHead()
			}
			if d.Tail.info == value{
				d.DeleteToTail()
			}
			return d
		}
		if node.info == value{
			// 删除node， 非头节点
			if node.prev != nil{
				node.next.prev = node.prev
				node.prev.next = node.next
				d.len --
			}else{
				// 头节点
				d.DeleteToHead()
			}
		}
		node = node.next

	}
}

// 通过index=>删除链表的节点
// index为正数 为从左->右 // 0表示第一个节点
// index为负数 为从右->左 // -1表示第一个节点
func (d *doubleList)DeleteToIndex(index int)*doubleList{
	// 链表为空
	if d.Head == nil{
		fmt.Println("链表为空")
		return d
	}
	// index 超过链表数
	if int(math.Abs(float64(index))) > d.len-1{
		fmt.Println("错误的index")
		return d
	}
	// 删除第一个node
	if index == 0{
		return d.DeleteToHead()
	}
	if index < 0{
		index += d.len
	}
	if index == d.len -1{
		return d.DeleteToTail()
	}
	_index := 1
	node := d.Head.next
	for {
		if index == _index{
			node.next.prev = node.prev
			node.prev.next = node.next
			d.len --
			return d
		}
		node = node.next
		_index ++
	}
}

// 遍历链表
func (d *doubleList) QuireAll() {
	if d.Head == nil{
		fmt.Println("链表数据为空")
		return
	}
	node := d.Head
	for {
		if node == nil{
			return
		}
		fmt.Println(node.info)
		if node == d.Tail{
			return
		}
		node = node.next
	}
}

// 判断valve是否存在
func (d *doubleList)QuireValue(value interface{}) bool{
	// 表为空
	if d.Head == nil{
		return false
	}
	// 表中只存在一个值
	if d.Head == d.Tail{
		if d.Head.info == value{
			return true
		}
		return false
	}
	// 遍历查值
	for node := d.Head; node != d.Tail.next; node = node.next{
		if node.info == value{
			return true
		}
	}
	return false
}

// 根据索引返回值
func (d *doubleList)QuireIndex(index int) interface{} {
	// 链表为空
	if d.Head == nil{
		return nil
	}
	if int(math.Abs(float64(index))) > d.len -1 {
		fmt.Println("索引值错误")
		return nil
	}
	if index == 0{
		return d.Head.info
	}
	if index == d.len -1 || index == -1{
		return d.Tail.info
	}
	if index < 0{
		index += d.len
	}
	__index := 1
	for node := d.Head.next; node != d.Tail; node = node.next{
		if __index == index{
			return node.info
		}
		__index ++
	}
	//fmt.Println("未能找到！！！！")
	return nil
}

func NewDoubleLinkedList()*doubleList{
	// 创建的链表头尾节点都为空
	return &doubleList{
		Head: nil,
		Tail: nil,
		len: 0,
	}
}



