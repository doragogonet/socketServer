package main

import (
	"fmt"
	"sync"

	"wshhz.com/Socket/SocketServer/SocketClient/Client"
)

var (
	wg sync.WaitGroup
)

const (
	socketServerAddress = "10.254.0.129:8090"
)

func init() {
	wg.Add(1)
}

func main() {
	Client.Start(&wg, socketServerAddress)

	wg.Wait()

	fmt.Println("SocketClient End...")
}
