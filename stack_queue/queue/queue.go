/*
author：zhaochengji mail：909536346@qq.com
基于双向链表的队列
*/
package queue

import (
	"github.com/zcj-git520/DataStructuresAlgorithmsForGo/Linked_list/double_linked_list"
)

type queueData struct {
	list *double_linked_list.DoubleList
}

// 获得队列的大小
func (q *queueData)Len() int{
	return q.list.Len()
}

// 将数据队尾
func (q *queueData)EnQueue(data interface{}) {
	q.list.AddToTail(data)
}

// 将数据从队首取出，并删除数据
func (q *queueData)DeQueue()interface{}{
	data := q.list.QuireIndex(0)
	// 栈不为空, 删除栈顶数据
	if data != nil{
		q.list.DeleteToHead()
	}
	return data
}

// 将数据取出，不删除数据
func (q *queueData)GetTopQueueValue()interface{}{
	return q.list.QuireIndex(0)
}

// 展示队列
func (q *queueData)ShowQueue() {
	q.list.QuireAll()
}

// 创建栈
func NewQueueData()*queueData{
	return  &queueData{list:double_linked_list.NewDoubleLinkedList()}
}