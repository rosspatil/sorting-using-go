package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(SortingMethod1([]float64{1.0, 3.0, 4.0, 2.0, 5.0}))
	fmt.Println(SortingMethod2([]float64{1.0, 3.0, 4.0, 2.0, 5.0}))
}

// SortingMethod1 -  this method sorts the array with extra space
func SortingMethod1(array []float64) []float64 {
	array1 := make([]float64, len(array))
	for i := 0; i < len(array); i++ {
		var min float64 = math.Inf(1)
		var pos = 0
		for j := 0; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				pos = j
			}
		}
		array[pos] = math.Inf(1)
		array1[i] = min
	}
	return array1
}

// SortingMethod2 - this method sorts the array without extra spaces
func SortingMethod2(array []float64) []float64 {
	for i := 0; i < len(array); i++ {
		var pos = 0
		var min = math.Inf(1)
		for j := i; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				pos = j
			}
		}
		tmp := array[i]
		array[i] = array[pos]
		array[pos] = tmp
	}
	return array
}
