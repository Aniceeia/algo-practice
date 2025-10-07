package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// context - используется для передачи контекста выполнения, таймайтов, дедлайнов,
// сигналов отмены между горутинами, это особенно важно для контроля над длительными
// операциями

func slowProcess(cnt context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Process ended")
		return nil
	case <-cnt.Done():
		fmt.Println("Process canceled")
		return cnt.Err()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	cnt := r.Context()
	err := slowProcess(cnt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
