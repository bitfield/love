package main

import (
	"fmt"
	"time"
)

func main() {
	copies := 1
	go func() {
		if copies > 0 {
			copies--
			fmt.Println("Customer A got the book")
		}
	}()
	go func() {
		if copies > 0 {
			copies--
			fmt.Println("Customer B got the book")
		}
	}()
	time.Sleep(time.Second)
}
