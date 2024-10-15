package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello World From Client")
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		log.Fatal(err)
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)

	fmt.Println("res", res)
}
