package main

import (
	"context"
	"fmt"
	"time"
)

func withTimeout(ctx context.Context) {
	cancelObserver := make(chan bool)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			cancelObserver <- true
			return
		}
	}()

	isCancel := <-cancelObserver
	if isCancel {
		close(cancelObserver)
		return
	}

	time.Sleep(10 * time.Second)
	fmt.Println("End of Program")
}

func setNameFromClient(ctx context.Context) {
	if value := ctx.Value("username"); value != nil {
		ctx := context.WithValue(ctx, "username_1", "Huy Tran")
		getListUsername(ctx)
	}
}

func getListUsername(ctx context.Context) {
	huyTran := ctx.Value("username_1")
	minhVo := ctx.Value("username")
	fmt.Println("Get List Username", huyTran, minhVo)
}

func main() {
	timeoutCtx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	valueCtx := context.WithValue(context.Background(), "username", "Minh Vo")
	fmt.Println("Starting..")

	withTimeout(timeoutCtx)
	setNameFromClient(valueCtx)
}
