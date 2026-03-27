package main

func main() {
	ch := make(chan int)

	close(ch)

	// Запись в закрытй канал - panic, так как канал уже закрыт
	ch <- 67
}
