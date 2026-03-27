package main

import "fmt"

func main() {
	ch := make(chan int)

	// В данном случае в val запишется значение из канала
	go func() {
		ch <- 67
	}()

	val := <-ch

	fmt.Println(val)
}
