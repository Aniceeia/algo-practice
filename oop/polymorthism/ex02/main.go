package main

import "fmt"

type iTax interface {
	calculate() int
}

type india struct {
	income        int
	taxPercentage int
}

func (i *india) calculate() int {
	return i.income * i.taxPercentage / 100
}

type bolivia struct {
	income        int
	taxPercentage int
}

func (b *bolivia) calculate() int {
	return b.income * b.taxPercentage / 100
}

func main() {
	i := &india{
		income:        12,
		taxPercentage: 22,
	}
	b := &bolivia{
		income:        11,
		taxPercentage: 33,
	}

	tax := []iTax{i, b}
	total := 0
	for _, v := range tax {
		total += v.calculate()
		fmt.Println(v.calculate())
	}
	fmt.Println(total)
}
