package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/minhvhd/go-server/logger"
)

func main() {

	http.HandleFunc("/hello", logger.Decorate(HelloHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	logger.Println(r.Context(), "HelloHandler Starting..")
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)

	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Timeout after 10 seconds")
	case <-ctx.Done():
		logger.Println(ctx, ctx.Err().Error())
	}
}
