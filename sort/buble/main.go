package main

import "fmt"

func bubleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		sorted := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				sorted = true
			}
		}
		if !sorted {
			break
		}
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{1, 4, 2, 5, 8, 11, 3, 6, 7, 10, 9}
	bubleSort(arr)
}
