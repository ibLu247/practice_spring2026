package main

func main() {
	ch := make(chan int)

	// В данном случае произойдет блокировка программы (deadlock)
	ch <- 67
}
