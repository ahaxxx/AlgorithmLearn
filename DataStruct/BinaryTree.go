package DataStruct

import "fmt"

type BinaryTreeNode struct {
	No int
	Name string
	Left *BinaryTreeNode
	Right *BinaryTreeNode
}

// 前序遍历
func PreOrder(node *BinaryTreeNode){
	if node != nil{
		fmt.Printf("no = %d name = %s\n",node.No,node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历
func InfixOrder(node *BinaryTreeNode){
	if node != nil{
		InfixOrder(node.Left)
		fmt.Printf("no = %d name = %s\n",node.No,node.Name)
		InfixOrder(node.Right)
	}
}

// 后序遍历
func PostOrder(node *BinaryTreeNode){
	if node != nil{
		PostOrder(node.Left)
		fmt.Printf("no = %d name = %s\n",node.No,node.Name)
		PostOrder(node.Right)
	}
}

func CreateBinaryTree(){
	root := &BinaryTreeNode{
		No:   1,
		Name: "宋江",
	}
	left1 := &BinaryTreeNode{
		No:   2,
		Name: "吴用",
	}
	node10 := &BinaryTreeNode{
		No:   10,
		Name: "李逵",
	}
	node12 := &BinaryTreeNode{
		No:   12,
		Name: "杨雄",
	}
	left1.Left = node10
	left1.Right = node12
	right1 := &BinaryTreeNode{
		No:   3,
		Name: "卢俊义",
	}
	root.Left = left1
	root.Right = right1
	right2 := &BinaryTreeNode{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2
	fmt.Println("前")
	PreOrder(root)
	fmt.Println("中")
	InfixOrder(root)
	fmt.Println("后")
	PostOrder(root)
}