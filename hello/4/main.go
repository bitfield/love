package main

import (
	"fmt"
	"time"
)

func goroutineB() {
	for i := range 10 {
		fmt.Println("Hello from goroutine B!", i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go goroutineB()
	for i := range 10 {
		fmt.Println("Hello from goroutine A!", i)
		time.Sleep(10 * time.Millisecond)
	}
}
