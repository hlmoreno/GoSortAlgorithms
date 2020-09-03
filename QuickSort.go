package main

import (
	"fmt"
)

func quickSort(entrada []int, low int, high int) {
	if low < high {
		index := ubicarPivot(entrada, low, high)
		quickSort(entrada, low, index-1)
		quickSort(entrada, index+1, high)
	}
}

func ubicarPivot(entrada []int, l int, h int) int {
	i := l
	pivot := entrada[h]

	for j := l; j < h; j++ {
		if entrada[j] <= pivot {
			entrada[j], entrada[i] = entrada[i], entrada[j]
			i++
		}
	}
	entrada[i], entrada[h] = entrada[h], entrada[i]
	fmt.Println(entrada)
	return i
}

func main() {
	vector := []int{10, 15, 13, -10, 5, 8, -3, 1, 20, 4}
	largo := len(vector)

	fmt.Println(vector)
	quickSort(vector, 0, largo-1)
	//fmt.Println(vector)
}
