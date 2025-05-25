package main

import (
	"fmt"
	"time"
)

func reader(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func writer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func doubler(ch <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for v := range ch {
			outCh <- v * 2
			time.Sleep(time.Second)
		}
		close(outCh)
	}()

	return outCh
}

func main3() {
	reader(doubler(writer()))
}
