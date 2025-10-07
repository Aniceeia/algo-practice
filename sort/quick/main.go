package main

import (
	"fmt"
)

func getMods() []map[int64]struct{} {
	mod := []int64{
		1e9 + 7,
		1e9 + 11,
		1e9 + 21,
	}

	fibNums := 40000

	modSet := make([]map[int64]struct{}, len(mod))
	for i := range mod {
		modSet[i] = make(map[int64]struct{})
	}

	//итерируемся по массиву и вписываем в мапу результат деления элементов массива фибоначчи

	for i, m := range mod {
		f1, f2 := int64(1), int64(1)

		modSet[i][1] = struct{}{}
		for j := 0; j < fibNums; j++ {
			f1, f2 = f2, (f1+f2)%m
			modSet[i][f2] = struct{}{}
		}

	}
	return modSet
}

func isFeb(number int64) bool {

}

func main() {

	fmt.Println(fibArr)
}
