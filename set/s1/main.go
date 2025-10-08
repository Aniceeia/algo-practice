package main

import (
	"bufio"
	"fmt"
	"os"
)

func getPisanoPeriods() ([]map[int64]struct{}, []int64) {
	mods := []int64{
		1e9 + 7,
		1e9 + 11,
		1e9 + 21,
	}
	pisanos := make([]map[int64]struct{}, len(mods))
	for p := range pisanos {
		pisanos[p] = make(map[int64]struct{})
	}
	fibMax := 40000
	for i, mod := range mods {
		f1, f2 := int64(1), int64(1)
		pisanos[i][1] = struct{}{}
		pisanos[i][1] = struct{}{}
		for range fibMax {
			f1, f2 = f2, (f1+f2)%mod
			pisanos[i][f2] = struct{}{}
		}
	}
	return pisanos, mods
}

func isFib(digits string, pisano []map[int64]struct{}, mods []int64) bool {
	for i, mod := range mods {
		var rem int64
		for _, ch := range digits {
			rem = (rem*10 + int64(ch-'0')) % mod
		}
		if _, exists := pisano[i][rem]; !exists {
			return false

		}
	}
	return true
}

func main() {
	pisano, mods := getPisanoPeriods()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	digits := scanner.Text()

	if isFib(digits, pisano, mods) {
		fmt.Println("yes")
		return
	} else {
		fmt.Println("no")
		return
	}

}
