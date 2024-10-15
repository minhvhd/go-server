package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello", HelloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloHandler Starting..")
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)

	defer cancel()

	select {
	//case <-time.After(10 * time.Second):
	//	fmt.Println("Timeout after 2 seconds")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
