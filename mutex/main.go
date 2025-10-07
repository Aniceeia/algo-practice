package main

import (
	"fmt"
	"sync"
)

// mutex - mutual exclusion - примитив синхронизации, позволяет толко одной горутине
// получить доступ к общему русурсу в один момент времени.

// методы
// Lock() - блокирует мьютекс. Если мьютекс уже заблокирован, горутина будет ждать
// пока он не освободиться
// Unlock() - разблокирует мьютекс. Вызывается после завершения работы с общим ресурсом
// обычно с помощью defer

func increment(mu *sync.Mutex, counter *int) {
	mu.Lock()
	defer mu.Unlock()
	*counter++
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	for range 387 {
		wg.Add(1)
		go func() {
			increment(&mu, &counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
