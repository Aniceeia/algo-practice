package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] < key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	fmt.Println(arr)
}

func main() {
	insertionSort([]int{1, 5, 3, 8, 88, 1383, 1})

}
