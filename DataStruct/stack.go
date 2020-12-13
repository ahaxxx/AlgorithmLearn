package DataStruct

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop 	int				// 最大栈顶
	Top		int				// 当前栈顶
	arr 	[5]int
}

// 入栈操作
func (this *Stack) Push(val int)(err error){
	if this.Top == this.MaxTop - 1{		// 判断栈是否已满
		fmt.Println("stack is fulled")
		return errors.New("stack is fulled")
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

// 出栈操作
func (this *Stack) Pop()(val int,err error){
	if this.Top == -1 {		// 判断栈是否为空
		fmt.Println("stack is empty !")
		return 0, errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val,nil
}

// 遍历栈,从栈顶开始遍历
func (this *Stack) List(){
	if this.Top == -1 {		// 判断栈是否为空
		fmt.Println("stack is empty !")
		return
	}
	for i:= this.Top;i >= 0; i--{
		fmt.Printf("stack[%d] is %d\n",i,this.arr[i])
	}
}

func CreatStack(){
	stack := Stack{
		MaxTop: 5,
		Top: -1,
	}
	stack.Push(5)
	stack.Push(9)
	stack.Push(8)
	stack.Push(9)
	stack.Push(52)
	stack.Push(2)
	stack.List()
	stack.Pop()
	stack.List()
	stack.Pop()
	stack.List()
}