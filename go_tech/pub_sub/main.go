package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond)
	defer cancel()

	apis := []string{"API1", "API2", "API3"}

	resultsCh := make(chan []string, len(apis))

	for _, api := range apis {
		go fetchFromAPI(ctx, api, resultsCh)
	}

	var aggregated []string
	for i := 0; i < len(apis); i++ {
		select {
		case res := <-resultsCh:
			aggregated = append(aggregated, res...)
		case <-ctx.Done():
			fmt.Println("Timeout reached while waiting for responses")
			break
		}
	}

	fmt.Println("Aggregated result:", aggregated)
}

func fetchFromAPI(ctx context.Context, api string, resultsCh chan<- []string) {
	apiCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	resCh := make(chan []string, 1)

	go func() {
		time.Sleep(time.Duration(500+rand.Intn(2500)) * time.Millisecond)

		resCh <- []string{api + "-result1", api + "-result2"}
	}()

	select {
	case <-apiCtx.Done():
		fmt.Println(api, "time out")
		return
	case res := <-resCh:
		resultsCh <- res
	}
}
