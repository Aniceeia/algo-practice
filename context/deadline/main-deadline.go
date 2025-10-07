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
		fmt.Println("process ended")
		return nil
	case <-cnt.Done():
		fmt.Println("process canceled")
		return cnt.Err()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	cnt := r.Context()
	deadline := time.Now().Add(1 * time.Second)
	cnx, cancel := context.WithDeadline(cnt, deadline)
	defer cancel()

	err := slowProcess(cnx)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "timeout exceeded "+deadline.Format(time.RFC1123), http.StatusGatewayTimeout)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "hello world")
	return
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
