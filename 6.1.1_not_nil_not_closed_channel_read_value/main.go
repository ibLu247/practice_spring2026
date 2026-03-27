package main

import "fmt"

func main() {
	ch := make(chan int)

	// В данном случае произойдет блокировка программы (deadlock)
	val := <-ch

	fmt.Println(val)
}
