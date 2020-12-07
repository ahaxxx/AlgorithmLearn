package DataStruct

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array	[5]int	//数组模拟队列
	head	int		//队首
	rear	int		//队尾
}

func (this *Queue) AddQueue(val int) (err error){
	//判断队列是否已满
	if this.rear == this.maxSize - 1{	// rear是队列尾部（含最后元素）
		return errors.New("Queue fulled")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

func (this *Queue) GetQueue() (val int,err error) {
	// 判断是否为空
	if this.rear == this.head{	// 队列
		return -1,errors.New("Queue empty")
	}
	this.head++
	val = this.array[this.head]
	return val,err
}

func (this *Queue) ShowQueue(){
	fmt.Println("队列当前情况为：")
	for i := this.head;i < this.rear;i++ {
		fmt.Printf("array[%d] = %d\t",i,this.array[i])
	}
}

func CreatQueue(){
	queue := &Queue{
		maxSize: 5,
		array: [5]int{},
		head: -1,
		rear: -1,
	}
	var key string
	var value int
	for{
		fmt.Println("\n\n1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出\n\n")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入入队的值")
			fmt.Scanln(&value)
			err := queue.AddQueue(value)
			if err != nil{
				fmt.Println(err.Error())
			}else {
				fmt.Print("加入队列ok")
			}
		case "get":
			value,err := queue.GetQueue()
			if err != nil{
				fmt.Println(err.Error())
			}else {
				fmt.Println("从队列中取出：",value)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}