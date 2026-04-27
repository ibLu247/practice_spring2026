package server

import (
	"bufio"
	"net"
	"sync"
)

var (
	mutex   sync.Mutex
	clients = make(map[net.Conn]struct{})
)

func handleConn(conn net.Conn) {
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
		conn.Close()
	}()

	mutex.Lock()
	clients[conn] = struct{}{}
	mutex.Unlock()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		broadcastMessage(message, conn)
	}
}

func broadcastMessage(message string, sender net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		if client != sender {
			client.Write([]byte(message + "\n"))
		}
	}
}
