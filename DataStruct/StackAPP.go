package DataStruct

import (
	"fmt"
	"strconv"
)

func (this *Stack) IsOper (val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47{
		return true
	}else {
		return false
	}
}

func (this *Stack) Cal (num1 int,num2 int,oper int) int{
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

func (this *Stack) Priority(oper int) int{
	res := 0
	if oper == 42 || oper == 47{
		res = 1
	}else if oper == 43 || oper == 45{
		res = 0
	}
	return res
}

func calculator(){
	numStack := &Stack{
		MaxTop: 20,
		Top:	-1,
	}
	opeStack := &Stack{
		MaxTop: 20,
		Top: 	-1,
	}
	exp := "30 + 3 * 6 - 4 -6"
	index 	:= 0
	num1 	:= 0
	num2	:= 0
	oper	:= 0
	result	:= 0
	keepNum	:= ""

	for{
		ch := exp[index : index + 1]
		temp := int([]byte(ch)[0])
		if opeStack.IsOper(temp){
			if opeStack.Top == -1{
				opeStack.Push(temp)
			} else {
				if opeStack.Priority(opeStack.arr[opeStack.Top]) >= opeStack.Priority(temp){
					num1,_ = numStack.Pop()
					num2,_ = numStack.Pop()
					oper,_ = opeStack.Pop()
					result = opeStack.Cal(num1,num2,oper)
					numStack.Push(result)
					opeStack.Push(temp)
				}else{
					opeStack.Push(temp)
				}
			}
		}else{
			keepNum += ch
			if index == len(exp) - 1{
				val,_ := strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			}else {
				if opeStack.IsOper(int([]byte(exp[index+1:index+2])[0])){

				}
			}
		}
	}
}