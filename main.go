package main

import (
	"com.leetcode/DataStruct"
	"fmt"
	"os"
)

func main()  {
	var key int
	for  {
		fmt.Println("请选择功能")
		fmt.Println("1.普通数组")
		fmt.Println("2.稀疏数组")
		fmt.Println("3.队列")
		fmt.Println("4.环形队列")
		fmt.Println("5.单向链表")
		fmt.Println("6.双向链表")
		fmt.Println("7.环形链表")
		fmt.Println("8.退出")
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
			os.Exit(0)
		default:
			fmt.Println("请检查输入后重新输入")
		}
	}
}