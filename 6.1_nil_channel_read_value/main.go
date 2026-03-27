package main

import "fmt"

func main() {
	var ch chan int

	// Чтение из nil канала (deadlock)
	val := <-ch

	// Эта часть программы никогда не выполнится
	fmt.Println(val)
}
