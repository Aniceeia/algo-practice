package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map
	sm.Store("hello", "world")
	sm.Store(1, 2)

	if v, ok := sm.Load("hello"); ok {
		fmt.Println("hello", v)
	}

	sm.Range(func(k, v any) bool {
		fmt.Println(k, v)
		return true
	}) //практиковать
}
