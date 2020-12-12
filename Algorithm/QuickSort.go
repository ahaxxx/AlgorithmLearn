package Algorithm

import "fmt"

func QuickSort(left int,right int,arr *[7]int){
	l := left
	r := right
	pivot := arr[(left+right)/2]
	for l<r {
		for arr[l] < pivot{
			r++
		}
		for arr[r] > pivot{
			l--
		}
		if l >=r {
			break
		}
		arr[l],arr[r] = arr[r],arr[l]
		if arr[l] == pivot{
			r--
		}
		if arr[r] == pivot {
			r++
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
	fmt.Printf("%v\n" ,*arr)
}