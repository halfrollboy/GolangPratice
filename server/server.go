package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type node struct {
	id   int
	conn net.Conn
}

var massNode [10]node

func getRandom() int {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 11000
	return rand.Intn((max - min + 1) + min)
}

func (n node) getSocketServer(i int) {
	ln, _ := net.Listen("tcp", ":808"+strconv.Itoa(i))
	n.conn, _ = ln.Accept()
}

func (n node) printAll() {
	fmt.Print(n.id)
	fmt.Println(n.conn)
}

func (n node) spam(wg *sync.WaitGroup) {
	// Запускаем цикл
	defer wg.Done()
	for {
		newmessage := strings.ToUpper(strconv.Itoa(getRandom()) + " It's " + strconv.Itoa(n.id))
		// Отправить новую строку обратно клиенту
		n.conn.Write([]byte(newmessage + "\n"))
	}
}

func main() {
	fmt.Println("Launching server...")

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		massNode[i] = node{id: i}
		massNode[i].getSocketServer(i)
	}

	for _, value := range massNode {
		wg.Add(1)
		go value.spam(&wg)
	}

	fmt.Println("progra, waiting")
	//Ожидание горутин
	wg.Wait()
}
