package DataStruct

import (
	"fmt"
)
/*
稀疏数组
案例：五子棋的存盘与复盘
*/
type ValueNode struct {
	row int 	//行
	col int		//列
	val int		//值
}

// 普通数组实现
func NormalArray(){
	var chessMap [11][11] int
	chessMap[1][2] = 1	//黑子
	chessMap[2][3] = 2	//白子

	//棋盘输出
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%d",v2)
		}
		fmt.Println()
	}
}

// 稀疏数组实现
func SpraseArray(){
	var chessMap [11][11] int
	chessMap[1][2] = 1	//黑子
	chessMap[2][3] = 2	//白子

	//切片
	var sparseArr []ValueNode
	//创建一个ValNode节点
	valueNode := ValueNode{
		row: 11,
		col: 11,
		val: 0,
	}
	//输入全局默认节点
	sparseArr = append(sparseArr,valueNode)
	for i,v := range chessMap{
		for j,v2 := range v{
			//创建一个节点
			if v2 != 0{
				valueNode := ValueNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr,valueNode)
			}
		}
	}
	//输出稀疏数组
	fmt.Println("当前稀疏数组：")
	for i,valueNode := range sparseArr{
		fmt.Printf("%d:%d %d %d\n",i,valueNode.row,valueNode.col,valueNode.val)
	}
	var chessMap2[11][11] int
	//稀疏数组恢复原始数组
	for i,valueNode := range sparseArr{
		//跳过第一行默认记录值
		if i!= 0{
			chessMap2[valueNode.row][valueNode.col] = valueNode.val
		}
	}
	for _,v :=range chessMap2{
		for _,v2 := range v{
			fmt.Printf("%d",v2)
		}
		fmt.Println()
	}
}