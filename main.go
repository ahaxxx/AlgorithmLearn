package main

import (
	"com.leetcode/Algorithm"
	"com.leetcode/DataStruct"
	"fmt"
	"os"
)

func main()  {
	var key int
	arr := [6]int{9,1,5,6,8,1}
	arr2 := [7]int{9,1,5,6,8,1,10}
	for  {
		fmt.Println("请选择功能")
		fmt.Println("1.普通数组")
		fmt.Println("2.稀疏数组")
		fmt.Println("3.队列")
		fmt.Println("4.环形队列")
		fmt.Println("5.单向链表")
		fmt.Println("6.双向链表")
		fmt.Println("7.环形链表")
		fmt.Println("8.环形链表应用")
		fmt.Println("9.选择排序")
		fmt.Println("10.插入排序")
		fmt.Println("11.快速排序")
		fmt.Println("12.栈")
		fmt.Println("13.栈应用")
		fmt.Println("14.哈希表")
		fmt.Println("15.退出")
		fmt.Println("请输入：")
		fmt.Scanln(&key)
		switch key {
		case 1:
			DataStruct.NormalArray()
		case 2:
			DataStruct.SpraseArray()
		case 3:
			DataStruct.CreatQueue()
		case 4:
			DataStruct.CreateCircleQueue()
		case 5:
			DataStruct.CreateHeroNode()
		case 6:
			DataStruct.CreatDoubleNode()
		case 7:
			DataStruct.CreateCricleNode()
		case 8:
			DataStruct.PlayGround()
		case 9:
			Algorithm.SelectSort(&arr)
		case 10:
			Algorithm.InsertSort(&arr2)
		case 11:
			Algorithm.UseQuickSort()
		case 12:
			DataStruct.CreatStack()
		case 13:
			DataStruct.Calculator()
		case 14:
			DataStruct.CreateHashTable()
		case 15:
			os.Exit(0)
		default:
			fmt.Println("请检查输入后重新输入")
		}
	}
}