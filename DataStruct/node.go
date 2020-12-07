package DataStruct

import "fmt"

type HeroNode struct {
	no int
	name string
	nickname string
	next *HeroNode
}

func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode)  {
	temp := head
	for  {
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	temp.next = newHeroNode
}

func InsertHeroNodeByNo(head *HeroNode,newHeroNode *HeroNode){
	temp := head
	flag := true
	for{
		if temp.next == nil{
			break
		}else if temp.next.no > newHeroNode.no{
			break
		}else if temp.next.no == newHeroNode.no{
			flag = false
			break
		}
		temp = temp.next
	}
	if !flag{
		fmt.Println("该英雄已经存在 no = ",newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func ListNode(head *HeroNode){
	temp := head
	if temp.next == nil{
		fmt.Println("链表是空的")
		return
	}
	for{
		fmt.Printf("[%d %s %s ==> \n]",temp.no,temp.name,temp.nickname)
		temp = temp.next
		if temp.next == nil {
			fmt.Printf("[%d %s %s \n]",temp.no,temp.name,temp.nickname)
			break
		}
	}
}

func DeleteNode(head *HeroNode,id int){
	temp := head
	flag := false
	for{
		if temp.next == nil{
			break
		}else if temp.next.no == id{
			fmt.Println("找到目标")
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	}else{
		fmt.Println("sorry,要删除的id,不存在")
	}
}

func CreateHeroNode(){
	head := &HeroNode{}
	hero1 := &HeroNode{
		no:1,
		name: "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}
	InsertHeroNode(head,hero1)
	ListNode(head)
	InsertHeroNode(head,hero2)
	ListNode(head)
	InsertHeroNode(head,hero3)
	ListNode(head)
	InsertHeroNode(head,hero4)
	ListNode(head)
	DeleteNode(head,1)
	ListNode(head)
	DeleteNode(head,2)
	ListNode(head)
	DeleteNode(head,3)
	ListNode(head)
	DeleteNode(head,4)
	ListNode(head)
	InsertHeroNodeByNo(head,hero4)
	ListNode(head)
	InsertHeroNodeByNo(head,hero3)
	ListNode(head)
	InsertHeroNodeByNo(head,hero2)
	ListNode(head)
	InsertHeroNodeByNo(head,hero1)
	ListNode(head)
}