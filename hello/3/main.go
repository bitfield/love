package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Println("Hello from goroutine A!", i)
	}
}
