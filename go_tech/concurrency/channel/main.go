package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	orders := generateOrders(20)

	//go func() {
	//	defer wg.Done()
	//	processOrders(orders)
	//}()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}

	// reportOrderStatus(orders)
	wg.Wait()
	fmt.Println("All operations compeleted.")
	fmt.Println(totalUpdates)
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

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateOrderStatus(order *Order) {
	order.mu.Lock()
	time.Sleep(
		time.Duration(rand.Intn(300)) *
			time.Millisecond,
	)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf(
		"Updated order %d\n status: %s",
		order.ID, order.Status,
	)
	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdate := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdate + 1
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
