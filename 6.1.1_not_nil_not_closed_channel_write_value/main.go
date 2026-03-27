package main

import "fmt"

func main() {
	ch := make(chan int)

	// А в этом случае блокировки не будет
	go func() {
		ch <- 67
	}()

	val := <-ch

	fmt.Println(val)
}
