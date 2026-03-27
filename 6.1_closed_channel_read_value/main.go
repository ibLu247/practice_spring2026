package main

import "fmt"

func main() {
	ch := make(chan int)

	close(ch)

	// Чтение из закрытого канала. В данном случае запишется дефолтное значение для типа int, т.е. 0
	val := <-ch

	fmt.Println(val)
}
