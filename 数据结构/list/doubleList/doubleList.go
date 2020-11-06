package main

import (
	"fmt"
	"sync"
)

//节点
type DoubleNode struct {
	Pre  *DoubleNode //前一个节点
	Next *DoubleNode //后一个节点
	Data interface{} //数据域
}

//双链表
type DoubleList struct {
	Head  *DoubleNode   //头节点
	Tail  *DoubleNode   //尾节点
	mutex *sync.RWMutex //读写锁
	Len   uint          //链表长度
}

//初始化双链表(我更习惯新建一个，可能有人习惯定义一个再初始化)
func (list *DoubleList) Init() {
	list.Head = nil
	list.Tail = nil
	list.Len = 0
	list.mutex = new(sync.RWMutex)
}

//添加节点到尾部
func (list *DoubleList) Append(newNode *DoubleNode) bool {
	if newNode == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Len == 0 {
		list.Head = newNode
		list.Tail = newNode
		newNode.Next = nil
		newNode.Pre = nil
		list.Len++
		return true
	}
	list.Tail.Next = newNode
	newNode.Pre = list.Tail
	newNode.Next = nil
	list.Tail = newNode
	list.Len++
	return true
}

//在指定节点后插入
func (list *DoubleList) Insert(index uint, newNode *DoubleNode) bool {
	if newNode == nil || index > list.Len {
		return false
	}
	if index == list.Len {
		return list.Append(newNode)
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Len == 0 { //插入头结点
		list.Head = newNode
		list.Tail = newNode
		newNode.Next = nil
		newNode.Pre = nil
		list.Len++
		return true
	}
	if index == 1 { //插入新的头结点
		newNode.Pre = nil
		newNode.Next = list.Head
		list.Head = newNode
		list.Len++
		return true
	}

	//这里用Get的话不行，读锁，写锁互斥
	//indexNode:=list.Get(index)
	var i uint
	indexNode := list.Head
	for i = 1; i < index; i++ {
		indexNode = indexNode.Next
	}

	newNode.Next = indexNode.Next
	newNode.Pre = indexNode
	newNode.Next.Pre = newNode
	indexNode.Next = newNode
	list.Len++
	return true
}

//删除节点
func (list *DoubleList) Delete(index uint) bool {
	if index > list.Len || index == 0 {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	//头节点，尾结点
	if index == 1 { //我觉得从1开始比较好，不想从0开始
		if list.Len == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head = list.Head.Next
			list.Head.Pre = nil
		}
		list.Len--
		return true
	}
	if index == list.Len { //尾结点
		list.Tail = list.Tail.Pre
		list.Tail.Next = nil
		list.Len--
		return true
	}

	//中间节点
	oldNode := list.Get(index)
	oldNode.Next.Pre = oldNode.Pre
	oldNode.Pre.Next = oldNode.Next
	list.Len--
	return true
}

//打印双链表
func (list *DoubleList) Display() {
	if list.Len == 0 || list == nil {
		fmt.Println("this double list id nil")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.Head
	index := 1
	for ptr != nil {
		fmt.Printf("%d, %v;", index, ptr.Data)
		ptr = ptr.Next
		index++
	}
}

//获取一个节点
func (list *DoubleList) Get(index uint) *DoubleNode {
	if index == 0 || index > list.Len {
		return list.Tail
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()
	var i uint
	tmp := list.Head
	for i = 1; i < index; i++ {
		tmp = tmp.Next
	}
	return tmp
}

func main() {
	var newDoubleList DoubleList
	newDoubleList.Init()
	for i := 1; i <= 10; i++ {
		newDoubleList.Append(&DoubleNode{
			Next: nil,
			Pre:  nil,
			Data: i,
		})
	}
	//newDoubleList.Delete(10)
	newDoubleList.Insert(9, &DoubleNode{
		Pre:  nil,
		Next: nil,
		Data: 99,
	})
	newDoubleList.Display()
}
