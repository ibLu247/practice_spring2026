package main

func main() {
	ch := make(chan int)

	close(ch)

	// Закрываем закрытый канал
	close(ch)
}
