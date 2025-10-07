package main

import "fmt"

// вариативные функции

type math struct {
}

func (m *math) add(n ...int) int {
	result := 0
	for i := range n {
		result += n[i]
	}
	return result
}

func main() {
	var m math

	r1 := m.add(1, 2, 3, 4)
	r2 := m.add(3, 4)

	fmt.Println(r1, r2)
}
