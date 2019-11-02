package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{1, 3, 4, 3, 2, 5, 8, 7, 6}, 0, 8))

}

// QuickSort - quick sort function
func QuickSort(array []int, start, end int) []int {
	if start < end {
		pIndex := Partition(array, start, end)
		QuickSort(array, start, pIndex-1)
		QuickSort(array, pIndex+1, end)
	}
	return array
}

// Partition - Partitioning function to break the array
func Partition(array []int, start, end int) int {
	pivot := array[end]
	pIndex := start
	for i := start; i < end; i++ {
		if array[i] < pivot {
			Swap(array, i, pIndex)
			pIndex++
		}
	}
	Swap(array, pIndex, end)
	return pIndex
}

// Swap - swap the element for given indices
func Swap(array []int, index1, index2 int) {
	tmp := array[index1]
	array[index1] = array[index2]
	array[index2] = tmp
}
