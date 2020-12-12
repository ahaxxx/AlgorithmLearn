package Algorithm

import "fmt"

func QuickSort(left int,right int,arr *[7]int){
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l<r {
		for arr[l] < pivot{
			l++
		}
		for arr[r] > pivot{
			r--
		}
		if l >=r {
			break
		}
		arr[l],arr[r] = arr[r],arr[l]
		if arr[l] == pivot{
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r{
		l++
		r--
	}
	if left < r {
		QuickSort(left,r,arr)
	}
	if left > l{
		QuickSort(l,right,arr)
	}
}

func UseQuickSort(){
	arr2 := [7]int{9,1,5,6,8,1,10}
	fmt.Println(arr2)
	QuickSort(0,len(arr2) - 1,&arr2)
	fmt.Println(arr2)
}