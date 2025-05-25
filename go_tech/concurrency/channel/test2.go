package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predicatbleWorker() {
	ch := make(chan int)
	go func() {
		randomTimeWork()
		close(ch)
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	select {
	case <-ch:
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("error")
	}
}

func main4() {
	predicatbleWorker()
}
