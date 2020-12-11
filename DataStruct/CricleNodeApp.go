package DataStruct

import (
	"fmt"
)

type Child struct {
	id int
	next *Child
}

func AddChild(num int) *Child{
	first := &Child{}
	curChild := &Child{}
	if num < 1{
		fmt.Println("非法值请重新输入")
		return first
	}
	for i := 1;i <= num; i++{
		child := &Child{
			id : i,
		}
		if i == 1 {
			first = child
			curChild = child
			curChild.next = first
		} else {
			curChild.next = child
			curChild = child
			curChild.next = first
		}
	}
	return first
}

func ShowChild(first *Child){
	if first.next == nil{
		fmt.Println("链表为空")
		return
	}

	curChild := first

	for{
		fmt.Printf("%d ==>\n",curChild.id)
		curChild = curChild.next
		if curChild.next == first {
			fmt.Printf("%d\n",curChild.id)
			break
		}
	}
}

/*
设编号为1，2，。。n的n个人未做一圈，约定编号为k(1<=k<=n)
的人从1开始报数，数到m的那个出列，它的下一位又从1开始报数。
数到m的那个人又出列，一次类推，直到所有人出列为止，由此产生一个出列编号的序列

*/
//分析思路
//1.编写一个函数，PlayGame(first *Boy,startNo int,countNum int)
//2.最后我们使用一个算法，按照要求，在环形链表中留下最后一个人

func Play(first *Child,startNo int,countNum int){
	//1.空的链表我们单独处理
	if first.next == nil{
		fmt.Println("链表为空")
		return
	}
	//留一个，判断startNo<=小孩总数
	//2.需要定义的辅助指针，帮助我们删除小孩
	tail := first
	//3.让tail执行环形链表的最后一个小孩，这个非常的重要
	for {
		if tail.next == first{
			break
		}
		tail = tail.next
	}

	for i := 1;i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	fmt.Println()
	for  {
		for i := 1;i <= countNum - 1;i++{
			first = first.next
			tail = tail.next
		}
		fmt.Printf("小孩编号为%d 出圈\n", first.id)
		first = first.next
		tail.next = first
		if tail == first {
			break
		}
	}
	fmt.Printf("小孩编号为%d出圈\n",first.id)
}

func PlayGround(){
	child := AddChild(10)
	ShowChild(child)
	Play(child,6,8)
}