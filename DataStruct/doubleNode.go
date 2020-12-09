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

func ListDoubleNodeNext(head *DoubleHeroNode){		// 从前往后遍历
	temp := head
	if temp.next == nil{
		fmt.Println("链表为空")
		return
	}
	for{
		fmt.Printf("[%d %s %s]==>\n",temp.no,temp.name,temp.nickname)
		temp = temp.next
		if temp.next == nil{
			fmt.Printf("[%d %s %s \n]",temp.no,temp.name,temp.nickname)
			break
		}
	}
}

func ListDoubleNodePre(head *DoubleHeroNode){		// 从后往前遍历
	temp := head
	if temp.next == nil{
		fmt.Println("链表为空")
		return
	}
	for{
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	for  {
		fmt.Printf("[%d %s %s]==>\n",temp.no,temp.name,temp.nickname)
		temp = temp.pre
		if temp.pre == nil{
			fmt.Printf("[%d %s %s \n]",temp.no,temp.name,temp.nickname)
			break
		}
	}
}

func CreatDoubleNode()  {
	head := &DoubleHeroNode{}
	hero1 := &DoubleHeroNode{
		no:1,
		name: "宋江",
		nickname: "及时雨",
	}

	hero2 := &DoubleHeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &DoubleHeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := &DoubleHeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	InsertDoubleHeroNode(head,hero1)
	InsertDoubleHeroNode(head,hero2)
	InsertDoubleHeroNode(head,hero3)
	InsertDoubleHeroNode(head,hero4)
	ListDoubleNodeNext(head)
	fmt.Println("逆序打印")
	ListDoubleNodePre(head)
	DeleteDoubleHeroNode(head,1)
	ListDoubleNodeNext(head)
	fmt.Println("逆序打印")
	ListDoubleNodePre(head)
}