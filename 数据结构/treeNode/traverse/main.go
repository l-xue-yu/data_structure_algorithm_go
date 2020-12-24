package main

import "fmt"

//参考博客：https://blog.csdn.net/tuobicui6522/article/details/80502130
//人家博客写的真好！

//二叉树结构
type treeNode struct {
	Value int
	left  *treeNode
	right *treeNode
}

//打印节点
func (node *treeNode) Print() {
	fmt.Print(node.Value, " ")
}

//添加
func CreateNode(data int) *treeNode {
	return &treeNode{Value: data}
}

//节点赋值
func (node *treeNode) SetValue(data int) {
	if node == nil { //空节点返回
		return
	}
}

//删除

//前序遍历： 根 前序左  前序右
func (node *treeNode) PreOrder() {
	if node == nil {
		return
	}

	node.Print()
	node.left.PreOrder()
	node.right.PreOrder()
}

//中序遍历：左 中 右
func (node *treeNode) middleOrder() {
	if node == nil {
		return
	}

	node.left.middleOrder()
	node.Print()
	node.right.middleOrder()
}

//后序遍历：左 右 中
func (node *treeNode) PostOrder() {
	if node == nil {
		return
	}

	node.left.PostOrder()
	node.right.PostOrder()
	node.Print()
}

//入口
func main() {
	root := treeNode{Value: 3} //根节点3
	root.left = &treeNode{}
	root.left.SetValue(0)           //左子树 第一个根节点 0
	root.left.right = CreateNode(2) //左子树 第二层
	root.right = &treeNode{5, nil, nil}
	root.right.left = CreateNode(4)

	fmt.Print("\n前序遍历：")
	root.PreOrder()
	fmt.Print("\n中序遍历：")
	root.middleOrder()
	fmt.Print("\n后序遍历：")
	root.PostOrder()

}
