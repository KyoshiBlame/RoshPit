package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("shutdown to stopped rating-worker")
			return
		default:
			time.Sleep(5 * time.Second)
			fmt.Println("rating worker working")
		}
	}
}
