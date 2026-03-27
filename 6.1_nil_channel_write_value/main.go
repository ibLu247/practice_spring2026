package main

import "fmt"

func main() {
	var ch chan int

	// Запись в nil канал
	ch <- 2

	// Блокировкак дальнейшего выполнения программы (deadlock)

	fmt.Println(ch)
}
