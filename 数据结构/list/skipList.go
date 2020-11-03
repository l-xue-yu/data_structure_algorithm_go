package main

import (
	"math/rand"
	"sync"
	"time"
)

//参考博客
//https://www.jianshu.com/p/400d24e9daa0?from=timeline&isappinstalled=0
//https://blog.csdn.net/xcl168/article/details/43205153?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-5.add_param_isCf&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-5.add_param_isCf
//https://classmatelin.blog.csdn.net/article/details/104324589?utm_medium=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.add_param_isCf&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-BlogCommendFromMachineLearnPai2-1.add_param_isCf
//https://www.jianshu.com/p/dc252b5efca6
type SkipListNode struct {
	key  int             //序号
	data interface{}     //数据
	next []*SkipListNode //节点指针切片，指向下一序号的指针数组
}

type SkipList struct {
	head   *SkipListNode //头节点
	tail   *SkipListNode //尾节点
	length int           //数据总量
	level  int           //层数
	mut    *sync.RWMutex //读写锁
	rand   *rand.Rand    //随机数生成器，用于生成层数
}

//生成随机层数
func (list *SkipList) randomLevel() int {

	level := 1

	for ; level < list.level && list.rand.Uint32()&0x1 == 1; level++ {
	}

	return level

}

func NewSkipList(level int) *SkipList {

	list := &SkipList{}

	if level <= 0 {

		level = 32

	}

	list.level = level
	list.head = &SkipListNode{next: make([]*SkipListNode, level, level)}

	list.tail = &SkipListNode{}

	list.mut = &sync.RWMutex{}

	list.rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	for index := range list.head.next {

		list.head.next[index] = list.tail

	}

	return list

}

func (list *SkipList) Add(key int, data interface{}) {

	list.mut.Lock()

	defer list.mut.Unlock()

	//1.确定插入深度

	level := list.randomLevel()

	//2.查找插入部位

	update := make([]*SkipListNode, level, level)

	node := list.head

	for index := level - 1; index >= 0; index-- {

		for {

			node1 := node.next[index]

			if node1 == list.tail || node1.key > key {

				update[index] = node //找到一个插入部位

				break

			} else if node1.key == key {

				node1.data = data

				return

			} else {

				node = node1

			}

		}

	}

	//3.执行插入

	newNode := &SkipListNode{key, data, make([]*SkipListNode, level, level)}

	for index, node := range update {

		node.next[index], newNode.next[index] = newNode, node.next[index]

	}

	list.length++

}

func (list *SkipList) Remove(key int) bool {

	list.mut.Lock()

	defer list.mut.Unlock()

	//1.查找删除节点

	node := list.head
	remove := make([]*SkipListNode, list.level, list.level)

	var target *SkipListNode

	for index := len(node.next) - 1; index >= 0; index-- {

		for {
			node1 := node.next[index]
			if node1 == list.tail || node1.key > key {
				break
			} else if node1.key == key {

				remove[index] = node //找到啦

				target = node1

				break

			} else {

				node = node1

			}

		}

	}

	//2.执行删除

	if target != nil {

		for index, node1 := range remove {

			if node1 != nil {

				node1.next[index] = target.next[index]

			}

		}

		list.length--

		return true

	}

	return false

}

func (list *SkipList) Find(key int) interface{} {

	list.mut.RLock()

	defer list.mut.RUnlock()

	node := list.head

	for index := len(node.next) - 1; index >= 0; index-- {

		for {

			node1 := node.next[index]

			if node1 == list.tail || node1.key > key {

				break

			} else if node1.key == key {

				return node1.data

			} else {

				node = node1

			}

		}

	}

	return nil

}

func (list *SkipList) Length() int {

	list.mut.RLock()

	defer list.mut.RUnlock()

	return list.length

}

func main() {
	newSkipList := NewSkipList(5)
	newSkipList.Add(1, 1)
	println("长度：", newSkipList.Length())
	newSkipList.Add(2, 1)
	println("长度：", newSkipList.Length())
}
