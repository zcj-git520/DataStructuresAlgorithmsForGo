package queue

import (
	"fmt"
	"testing"
)

func TestNewQueueData(t *testing.T) {
	q := NewQueueData()
	q.EnQueue(0)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.ShowQueue()
	fmt.Println("*************************************")
	data := q.GetTopQueueValue()
	fmt.Println("data = >", data)
	data = q.GetTopQueueValue()
	fmt.Println("data = >", data)
	data =  q.GetTopQueueValue()
	fmt.Println("data = >", data)
	q.ShowQueue()
	fmt.Println("************************************")
	data = q.DeQueue()
	fmt.Println("data = >", data)
	q.ShowQueue()
	fmt.Println("************************************")
	data = q.DeQueue()
	fmt.Println("data = >", data)
	q.ShowQueue()
	fmt.Println("************************************")
	data = q.DeQueue()
	fmt.Println("data = >", data)
	q.ShowQueue()
	fmt.Println("************************************")
	fmt.Println("stack len is ", q.Len())

}
