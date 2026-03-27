package main

func main() {
	ch := make(chan int)

	// Просто закрываем канал
	close(ch)
}
