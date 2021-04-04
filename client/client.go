package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type node struct {
	id   int
	conn net.Conn
}

var massNode [10]node

func (n node) getSocketServer(i int) {
	n.conn, _ = net.Dial("tcp", "127.0.0.1:808"+strconv.Itoa(i))
}

func (n node) connection(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Прослушиваем ответ
		message, _ := bufio.NewReader(n.conn).ReadString('\n')
		fmt.Print("Message: " + message)
	}
}

func main() {
	fmt.Println("Client launch")
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		massNode[i] = node{id: i}
		massNode[i].getSocketServer(i)
	}

	for _, value := range massNode {
		wg.Add(1)
		go value.connection(&wg)
	}

	time.Sleep(time.Millisecond)
	// Ожидание горутин
	wg.Wait()
}
