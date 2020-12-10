package DataStruct

import "fmt"

type CricleNode struct {
	no int
	name string
	next *CricleNode
}

func InsertCricleNode(head *CricleNode,newCricleNode *CricleNode){
	if head.next == nil{
		head.no = newCricleNode.no
		head.name = newCricleNode.name
		head.next = head
		fmt.Println(newCricleNode,"已加入")
		return
	}

	temp := head
	for  {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	temp.next = newCricleNode
	newCricleNode.next = head
}

func ListCircleNode(head *CricleNode){
	if head.next == nil{
		println("链表为空")
		return
	}

	temp := head

	for  {
		fmt.Printf("%d %s ==>\n",temp.no,temp.next)
		temp = temp.next
		if temp.next == head {
			fmt.Printf("%d %s \n",temp.no,temp.next)
			break
		}
	}
}

func DeleteCircleNode(head *CricleNode,id int) *CricleNode{
	temp := head
	helper := head
	if temp.next == nil{
		fmt.Println("链表为空")
		return head
	}

	if temp.next == head{
		if temp.no == id {
			temp.next = nil
			fmt.Printf("删除节点%d\n",id)
		}
		return head
	}
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true

	for{
		if temp.next == head {
			break
		}
		if temp.no == id{
			if temp == head{
				head = head.next
			}
			helper.next = temp.next
			fmt.Printf("删除节点%d\n",id)
			flag = false
			break
		}
		temp = temp.next
		helper = helper.next
	}
	if flag{
		fmt.Println("未找到节点%d",id)
	}
	return head
}

func CreateCricleNode(){
	head := &CricleNode{}
	//创建一只猫
	cat1 := &CricleNode{
		no:   1,
		name: "tom1",
	}
	cat2 := &CricleNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CricleNode{
		no:   3,
		name: "tom3",
	}
	cat4 := &CricleNode{
		no:   4,
		name: "tom4",
	}
	cat5 := &CricleNode{
		no:   5,
		name: "tom5",
	}
	InsertCricleNode(head,cat1)
	InsertCricleNode(head,cat2)
	InsertCricleNode(head,cat3)
	InsertCricleNode(head,cat4)
	InsertCricleNode(head,cat5)
	ListCircleNode(head)
	head = DeleteCircleNode(head, 1)
	ListCircleNode(head)
	head = DeleteCircleNode(head, 4)
	ListCircleNode(head)
}