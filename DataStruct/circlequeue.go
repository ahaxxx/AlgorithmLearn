package DataStruct

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int		// 最大长度
	array	[5]int	// 数据数组
	head	int		//队首
	tail	int		//队尾
}

// 入队
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull(){
		return errors.New("Queue is Fulled")
	}
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return
}

// 出队
func (this *CircleQueue) Pop() (val int,err error){
	if this.IsEmpty() {
		return 0,errors.New("Queue is empty")
	}
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// 显示队列
func (this *CircleQueue) ListQueue(){
	fmt.Println("队列如下：")
	size := this.Size()
	if size == 0 {
		fmt.Println("Queue is empty")
	}
	tempHead := this.head
	for i := 0;i < size;i++{
		fmt.Printf("arr[%d] = %d\t",tempHead,this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 队满判断
func (this *CircleQueue) IsFull() bool {
	if (this.tail + 1) % this.maxSize == this.head {
		return true
	} else{
		return false
	}
}

// 队空判断
func (this *CircleQueue) IsEmpty() bool{
	if this.tail == this.head {
		return true
	} else {
		return false
	}
}

// 队列长度获取
func (this *CircleQueue) Size() int{
	size := len(this.array)
	return size
}

func CreateCircleQueue(){
	cq := CircleQueue{
		maxSize: 5,
		array: [5]int{},
		head: 0,
		tail: 0,
	}
	var key string
	var val int
	for{
		fmt.Println("\n\n1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出\n\n")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入入队值：")
			fmt.Scanln(&val)
			err := cq.Push(val)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("入队成功")
			}
		case "get":
			val,err := cq.Pop()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("出队成功：" ,val)
			}
		case "show":
			cq.ListQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}