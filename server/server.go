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

func getSocketServer(i int) net.Conn {
	ln, _ := net.Listen("tcp", ":808"+strconv.Itoa(i))
	conn, _ := ln.Accept()
	return conn
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

func (c node) clientRequest(conn net.Conn) {
	defer conn.Close() //Закрываем соединение по выходу из функции

	buf := make([]byte, 32) //Буфер для чтения клиентских данных
	for {
		conn.Write([]byte("This clientRequest = " + strconv.Itoa(c.id))) //пишем в сокет

		readLen, err := conn.Read(buf) //Читаем из сокета
		if err != nil {
			fmt.Println(err)
			break
		}
		conn.Write(append([]byte("hello"), buf[:readLen]...)) //Пишем в сокет
	}
}

func main() {
	fmt.Println("Launching server...")

	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		fmt.Println(i)
		massNode[i] = node{id: i, conn: getSocketServer(i)}

	}

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go massNode[i].spam(&wg)
	}

	fmt.Println("progra, waiting")
	//Ожидание горутин
	wg.Wait()
}
