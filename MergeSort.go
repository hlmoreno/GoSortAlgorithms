package main

import "fmt"

func mergeSort(entrada []int, low int, high int) {
	var mid int
	if low < high {
		mid = (low + high) / 2
		mergeSort(entrada, low, mid)
		mergeSort(entrada, mid+1, high)
		merge(entrada, low, mid, high)
	}
}

func merge(entrada []int, l int, m int, h int) {
	//arr1 := entrada[l:m]
	//arr2 := entrada[m:h]

	arr1 := make([]int, m-l+1)
	arr2 := make([]int, h-m)

	copy(arr1, entrada[l:m+1])
	copy(arr2, entrada[m+1:h+1])

	if len(arr1) == 0 || len(arr2) == 0 {
		return
	}

	i := 0
	j := 0
	for k := l; k <= h; k++ {
		if len(arr1) > i && len(arr2) > j {
			if arr1[i] <= arr2[j] {
				entrada[k] = arr1[i]
				i++
			} else {
				entrada[k] = arr2[j]
				j++
			}
		} else if len(arr1) > i {
			entrada[k] = arr1[i]
			i++
		} else {
			entrada[k] = arr2[j]
			j++
		}

	}
	fmt.Println(entrada)
}

func main() {
	vector := []int{10, 15, 13, -10, 5, 8, -3, 1, 20, 4}
	largo := len(vector)

	fmt.Println(vector)
	mergeSort(vector, 0, largo-1)
	fmt.Println(vector)
}
