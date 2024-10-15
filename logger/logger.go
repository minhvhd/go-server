package logger

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
)

type key int64

const requestIDKey = key(1)

func Println(ctx context.Context, msg string) {
	id, ok := ctx.Value(requestIDKey).(int64)
	if !ok {
		fmt.Println("Could not get request ID")
		return
	}

	fmt.Printf("[%d] %s\n", id, msg)
}

func Decorate(callBack http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := rand.Int63()
		ctx := context.WithValue(r.Context(), requestIDKey, id)

		callBack.ServeHTTP(w, r.WithContext(ctx))
	})
}
