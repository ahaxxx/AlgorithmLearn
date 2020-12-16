package DataStruct

import (
	"fmt"
	"os"
)

type Emp struct {
	Id		int
	Name 	string
	Next	*Emp
}

type EmpLink struct {
	Head *Emp
}

type HashTable struct {
	LinkArr [7]EmpLink
}

func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp) //
}

func (this *Emp) ShowMe(){
	fmt.Printf("链表%d 找到该雇员 %d\n", this.Id%7, this.Id)
}

// 添加员工
func (this *EmpLink) Insert(emp *Emp){
	cur := this.Head
	var pre *Emp = nil
	// 如果当前的EmpLink就是一个空链表
	// 为空的话直接放到开头
	if cur == nil{
		this.Head = emp
		return
	}
	for{
		if cur != nil{
			if cur.Id > emp.Id{
				break
			}
			pre = cur
			cur = cur.Next
		} else{
			break
		}
	}
	pre.Next = emp
	emp.Next = cur
}

func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil{
		fmt.Printf("链表%d为空\n",no)
		return
	}
	cur := this.Head
	for{
		if cur != nil{
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else{
			break
		}
	}
	fmt.Println("")
}

func (this *EmpLink) FindById(id int) *Emp{
	cur := this.Head
	for {
		if cur != nil && cur.Id == id{
			return cur
		}else if cur == nil{
			break
		}
		cur = cur.Next
	}
	return nil
}

func (this *HashTable) ShowAll(){
	for i := 0;i < len(this.LinkArr);i++{
		this.LinkArr[i].ShowLink(i)
	}
}

func (this *HashTable) HashFun(id int) int{
	return id % 7
}

func (this *HashTable) FindById(id int) *Emp{
	LinkNo := this.HashFun(id)
	return this.LinkArr[LinkNo].FindById(id)
}

func CreateHashTable(){
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for{
		fmt.Println("===============雇员系统菜单============")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号:")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//编写一个方法，显示雇员信息
				emp.ShowMe()
			}

		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}