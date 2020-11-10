package main

import (
	"fmt"
	"sync"
)

//节点
type CycleNode struct {
	data interface{}
	next *CycleNode
}

//单向循环链表
type CycleList struct {
	len   uint
	head  *CycleNode
	tail  *CycleNode
	mutex *sync.RWMutex //读写锁
}

//初始化
func (cList *CycleList) Init() {
	cList.head = nil
	cList.tail = nil
	cList.len = 0
	cList.mutex = new(sync.RWMutex)
}

//尾部添加节点
func (clist *CycleList) Append(data interface{}) bool {
	node := new(CycleNode)
	(*node).data = data

	clist.mutex.Lock()
	defer clist.mutex.Unlock()
	if clist.GetSize() == 0 { //单独处理头结点
		clist.head = node
		clist.tail = node
	} else {
		if clist.len == 1 { //只有一个头结点
			clist.head.next = node
			clist.tail = node
		} else {
			tmp := clist.head
			for ; tmp.next != clist.head; tmp = tmp.next { //寻找尾节点

			}
			tmp.next = node //新节点放到尾节点后面
		}

	}
	node.next = clist.head //尾节点指向头结点
	clist.len++
	return true
}

//在当前节点的尾部插入新的节点
func (clist *CycleList) Insert(elementNode *CycleNode, data interface{}) bool {
	if elementNode == nil {
		return false
	}

	node := new(CycleNode)
	node.data = data

	clist.mutex.Lock()
	defer clist.mutex.Unlock()

	node.next = elementNode.next //待插入节点指向前节点的下一个节点
	elementNode.next = node      //前节点，指向待插入节点

	clist.len++
	return true
}

//删除节点
func (cList *CycleList) Delete(elementNode *CycleNode) bool {
	if elementNode == nil {
		return false
	}

	cList.mutex.Lock()
	defer cList.mutex.Unlock()

	//头节点
	if elementNode == cList.head {
		if cList.len == 1 {
			cList.head = nil
			cList.tail = nil
			cList.len = 0
		} else {
			cList.head = cList.head.next
			cList.tail.next = cList.head
			cList.len--
		}
	}
	tmp := cList.head
	for ; tmp.next != elementNode; tmp = tmp.next {

	}
	tmp.next = elementNode.next
	cList.len--
	return true
}

//获取链表开头

//获取链表节点数量
func (cList *CycleList) GetSize() uint {
	return cList.len
}

//获取一个节点
func (cList *CycleList) GetNode(index uint) *CycleNode {
	if index == 0 || index > cList.len {
		return cList.head //返回尾节点比较好
	}

	cList.mutex.RLock()
	defer cList.mutex.RUnlock()

	tmp := cList.head
	var i uint
	for i = 1; i < index; i++ {
		tmp = tmp.next
	}
	return tmp
}

//打印循环链表
func (cList *CycleList) Delay() {
	tmp := cList.head
	index := 1
	for ; tmp.next != cList.head; tmp = tmp.next {
		println("")
		fmt.Printf("%d,%v,", index, tmp)
		index++
	}
	println("")
	fmt.Printf("%d,%v,", index, tmp)

}

func main() {
	var newCycleList CycleList
	newCycleList.Init()
	newCycleList.Append(1)
	newCycleList.Delay()
	println("len:", newCycleList.len)
	newCycleList.Insert(newCycleList.GetNode(1), 2)
	newCycleList.Delay()
	println("len:", newCycleList.len)
	newCycleList.Delete(newCycleList.GetNode(2))
	newCycleList.Delay()
}
