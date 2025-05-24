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

	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}

		close(orderChan)

		fmt.Println("Done")
	}()

	go processOrders(orderChan, &wg)

	// reportOrderStatus(orders)
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

func processOrders(orderChan <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func reportOrderStatus(orders []*Order) {
	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %s\n",
			order.ID, order.Status,
		)
	}
	fmt.Println("\n---------------------------")
}
