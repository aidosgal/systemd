package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	orderChan := make(chan *Order, 20)
	processedChan := make(chan *Order, 20)

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}

		close(orderChan)

		fmt.Println("Done")
	}()

	go processOrders(orderChan, processedChan, &wg)

	go func() {
		defer wg.Done()
		for {
			select {
			case processedOrder, ok := <-processedChan:
				if !ok {
					fmt.Println("processing channel closed")
					return
				}
				fmt.Println(processedOrder)
			case <-time.After(time.Minute):
				fmt.Println("time")
			}
		}
	}()

	wg.Wait()
	fmt.Println("All operations compeleted.")
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}

	return orders
}

func processOrders(
	inChan <-chan *Order,
	outChan chan<- *Order,
	wg *sync.WaitGroup,
) {
	defer func() {
		wg.Done()
		close(outChan)
	}()
	for order := range inChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		order.Status = "Processing"
		outChan <- order
	}
}
