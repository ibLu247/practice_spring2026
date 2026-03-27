package main

func main() {
	var ch chan int

	// Закрытие nil канала
	close(ch)
}
