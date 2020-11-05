package main

import (
	"fmt"
	"sync"
)

//参考博客
//https://www.jianshu.com/p/cdc223a289e4
//使用读写锁保证并发安全

//节点
type singleNode struct {
	data interface{} //数据域
	next *singleNode //指针域
}

//单链表
type SingleList struct {
	Head  *singleNode   //头节点
	Tail  *singleNode   //尾节点
	Len   uint          //链表长度
	mutex *sync.RWMutex //读写锁
}

//创建新链表
func NewSingleList() SingleList {
	//一个链表一个读写锁
	newSingleList := SingleList{
		Head:  nil,
		Tail:  nil,
		Len:   0,
		mutex: new(sync.RWMutex),
	}
	return newSingleList
}

//添加节点到尾部
func (list *SingleList) Append(node *singleNode) bool {
	//空节点
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()

	//如果是空链表,头结点指向node,尾结点指向node
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
		list.Len = 1
		return true
	}
	//正常尾部添加
	list.Tail.next = node
	list.Tail = node
	list.Len++
	return true
}

//在头部添加节点
func (list *SingleList) add(node *singleNode) bool {
	if node == nil {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Len == 0 {
		list.Head = node
		list.Head = node
		list.Len = 1
		return true
	}

	//缓存原头结点
	node.next = list.Head
	list.Head = node
	list.Len++
	return true
}

//在指定位置插入节点，在第index位置后面插入
func (list *SingleList) insert(index uint, node *singleNode) bool {
	if node == nil {
		return false
	}
	//超过链表的长度
	if index > list.Len {
		return false
	}
	if list.Len == 0 {
		list.Head = node
		list.Tail = node
		list.Len = 1
		return true
	}
	var i uint
	ptr := list.Head
	for i = 1; i < index; i++ {
		ptr = ptr.next
	}
	//此时已经是待插入节点
	node.next = ptr.next
	ptr.next = node
	list.Len++
	return true
}

//删除节点，删除第index个节点
func (list *SingleList) remove(index uint) bool {
	if index > list.Len {
		return false
	}

	//头节点删除
	if index == 1 {
		if list.Len == 1 { //只有一个节点
			list.Head = nil
			list.Tail = nil
			list.Len = 0
			return true
		}
		list.Head = list.Head.next
		list.Len--
		return true
	}

	var i uint
	ptr := list.Head
	beforeNode := list.Head
	for i = 1; i < index; i++ {
		beforeNode = ptr //前面一个节点
		ptr = ptr.next   //ptr现在是要删除的节点
	}
	//尾节点删除
	if index == list.Len {
		list.Tail = beforeNode
		beforeNode.next = nil
		list.Len--
		return true
	}
	beforeNode.next = ptr.next
	list.Len--
	return true

}

//获取节点信息，获取第index节点
func (list *SingleList) get(index uint) *singleNode {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	var i uint
	ptr := list.Head
	for i = 1; i < index; i++ {
		ptr = ptr.next
	}
	return ptr
}

//打印链表
// 输出链表
func (list *SingleList) Display() {
	if list == nil {
		fmt.Println("this single list is nil")
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	fmt.Printf("this single list size is %d \n", list.Len)
	ptr := list.Head
	var i uint
	for i = 0; i < list.Len; i++ {
		fmt.Printf("No%3d data is %v\n", i+1, ptr.data)
		ptr = ptr.next
	}
}

func main() {
	newList := NewSingleList()

	//插入100个
	for i := 1; i <= 100; i++ {
		newList.Append(&singleNode{
			data: i,
			next: nil,
		})
	}
	//在第50个节点后插入新节点
	newList.insert(50, &singleNode{
		data: 555,
		next: nil,
	})
	newList.Display()
	//删除最后一个节点
	newList.remove(51)
	newList.Display()

	newList.get(100)
	fmt.Printf("No1 data is %v\n", newList.get(1).data)
	fmt.Printf("No100 data is %v\n", newList.get(100).data)
}
