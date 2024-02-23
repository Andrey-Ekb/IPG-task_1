package main

import (
	"fmt"
)

func sort(array [9]int) [9]int {
	per := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				per = array[j]
				array[j] = array[j+1]
				array[j+1] = per
			}
		}
	}
	return array
}

func main() {
	array := []int{0, -5, 6, -1}
	array1 := []int{-20, 1, 40, 20, -9}

	fmt.Println()
	fmt.Println(array)
	fmt.Println(array1)
	fmt.Println()

	for k := 0; k < len(array1); k++ {
		array = append(array, array1[k])
	}

	fmt.Printf("Неотсортированный массив %d\n", array)
	sortarray := sort([9]int(array))
	fmt.Printf("Отсортированный массив   %d\n", sortarray)
	fmt.Println()
}
