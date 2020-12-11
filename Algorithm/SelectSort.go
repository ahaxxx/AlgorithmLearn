package Algorithm

import "fmt"

func SelectSort(arr *[6]int){
	for j := 0;j < len(arr) - 1;j++{
		max := arr[j]
		maxIndex := j
		for i := j + 1;i < len(arr); i++{
			if max < arr[i]{
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j{
			arr[j],arr[maxIndex] = arr[maxIndex],arr[j]
		}
		fmt.Printf("第%d次%v\n", j+1, *arr)
	}
}