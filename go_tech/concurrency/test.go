package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID int
}

func ProcessOrders(orders []Order, workerCount int) {
	jobs := make(chan Order)
	wg := &sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for order := range jobs {
				process(order)
			}
		}()
	}

	for _, order := range orders {
		jobs <- order
	}
	close(jobs)

	wg.Wait()
}

func process(o Order) {
	fmt.Printf("Processing order %d\n", o.ID)
	time.Sleep(time.Second)
}
