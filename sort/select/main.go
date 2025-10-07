package main

import "fmt"

func selectSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println(arr)
}

func main() {
	arr := []int{1, 4, 2, 5, 8, 11, 3, 6, 7, 10, 9}
	selectSort(arr)
}
