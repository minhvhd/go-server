package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello World From Client")
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	if err != nil {
		fmt.Println("Error creating request", err)
	}

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error making request", err)
	}

	defer res.Body.Close()

	fmt.Println("res", res.StatusCode)

}
