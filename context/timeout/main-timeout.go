package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func handlerDeadline(w http.ResponseWriter, r *http.Request) {
	cnt := r.Context()
	ctx, cancel := context.WithTimeout(cnt, 1*time.Second)
	defer cancel()
	err := slowProcess(ctx)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "deadline exceeded", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "hello, world")
}

func main() {
	http.HandleFunc("/", handlerDeadline)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
