package Algorithm

import "fmt"

func InsertSort(arr *[7]int){
	for i := 1; i < len(arr); i++{
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal{
			fmt.Printf("arr[%v] = %v < %v\n",insertIndex,arr[insertIndex],insertVal)
			arr[insertIndex + 1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex + 1 != i{
			fmt.Printf("insertIndex = %v\n",insertIndex)
			arr[insertIndex + 1] = insertVal
		}
		fmt.Printf("第%d次插入后 %v\n", i, *arr)
	}
}