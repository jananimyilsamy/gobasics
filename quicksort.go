
package main


import (
	"fmt"
)

//Quick Sorter method
func QuickSort(arr []int, low int, high int) {
	if low < high {
		var part int
		part = partition(arr, low, high)
		QuickSort(arr, low, part-1)
		QuickSort(arr, part+1, high)
	}
}


func partition(arr []int, low int, high int) int {
	var pivot int
	pivot = arr[high]
	var i int
	i = low
	var j int
	for j = low; j < high; j++ {
		if arr[j] <= pivot {
			swap(&arr[i], &arr[j])
			i += 1
		}
	}
	swap(&arr[i], &arr[high])
	return i
}


func swap(t1 *int, t2 *int) {
	var val int
	val = *t1
	*t1 = *t2
	*t2 = val
}


func main() {
	var num int

	fmt.Print("Enter size: ")
	fmt.Scan(&num)

	var array = make([]int, num)

	var i int
	for i = 0; i < num; i++ {
		
		fmt.Scan(&array[i])
	}

	
	QuickSort(array, 0, num-1)
	fmt.Print("Sorted Elements: ", array, "\n")
}
