package DataStruct

import "fmt"

type DoubleHeroNode struct {
	no int
	name string
	nickname string
	pre *DoubleHeroNode			// 指向前一节点
	next *DoubleHeroNode		// 指向后一节点
}

func InsertDoubleHeroNode(head *DoubleHeroNode,newHeroNode *DoubleHeroNode){
	temp := head			// 创建辅助结点
	for {					// 寻找最后一个节点
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode	// 插入尾结点
	newHeroNode.pre = temp	// 链接头部节点
}

func InsertDoubleHeroNodeById(head *DoubleHeroNode,id int, newHeroNode *DoubleHeroNode){
	temp := head
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no > newHeroNode.no{
			break
		} else if temp.next.no == newHeroNode.no{
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag {
		fmt.Println("对不起，该英雄已经存在 no = ",newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil{
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

func DeleteDoubleHeroNode(head *DoubleHeroNode,id int){
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.no == id{
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
		if temp.next != nil{
			temp.next.pre = temp
		}
	} else {
		fmt.Println("节点不存在")
	}
}