package main

import "fmt"

func cambiarPos(num1 *int, num2 *int) {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}

func bubbleSort(entrada []int, largo int) {
	for i := 0; i < largo-1; i++ {
		cambio := false
		for j := 0; j < largo-i-1; j++ {
			if entrada[j] > entrada[j+1] {
				cambiarPos(&entrada[j], &entrada[j+1])
				cambio = true
			}
		}
		if !cambio {
			break
		}
		fmt.Println(entrada)
	}
}

func main() {
	vector := []int{10, 15, 13, -10, 5, 8, -3, 1, 20, 4}
	largo := len(vector)
	fmt.Println(vector)
	bubbleSort(vector, largo)
	//fmt.Println(vector)
}
