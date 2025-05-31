package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	copies := 1
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if copies > 0 {
			copies--
			fmt.Println("Customer A got the book")
		}
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		if copies > 0 {
			copies--
			fmt.Println("Customer B got the book")
		}
	}()
	time.Sleep(time.Second)
}
