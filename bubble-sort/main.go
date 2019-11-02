package main

import "fmt"

func main() {
	fmt.Println(Sorting([]int{1, 3, 4, 2, 5}))
}

// Sorting - this sorting algo sorts array using bubble sort
func Sorting(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		swap := false
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return array
}
