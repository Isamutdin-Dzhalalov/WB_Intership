package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

const address = "localhost:8080"

func main() {
	go runServer()
	// Парсинг аргументов командной строки.
	timeOut := flag.Duration("timeout", 10*time.Second, "Time out flag")
	flag.Parse()
	// Задержка, чтобы сервер успел запуститься.
	time.Sleep(time.Second)
	/* Ф-ция для ассинхронного подключения к удалённому серверу.
	   В аргументах: протокол для соединения, адресс, время ожидания подключения
	   к серверу, при неудачном подключении вернёт ошибку.*/
	conn, err := net.DialTimeout("tcp", address, *timeOut)
	if err != nil {
		log.Fatalln(err)
	}
	go writeToSocket(conn)

	for {
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed")
			break
		}
		fmt.Println("Get message from the socket: ", text)
	}

}

func writeToSocket(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
	/* Считывает ввод с клавиатуры, пока не нажмём "Enter".
	   err == io.EOF если мы нажмём "Ctlr+D".*/
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Got ctrl+D signal.Closing connection")
			// Закрытие соединения.
			conn.Close()
			return
		}
		_, err = conn.Write(text)
		if err != nil {
			conn.Close()
			fmt.Println("Connection closed due to error:", err)
			return
		}
	}
}
func runServer() {
	// tcp-слушатель указанного адресса.
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	// Ожидаем подключения. В случае неудачи, возвращает ошибку.
	conn, err := listener.Accept()
	if err != nil{
		return
	}
	// Перед выходом из функции закрываем соединение.
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	for {
	/* Читаем строку из соединения conn до тех пор, пока не
	   встретим '\n'. 
	   Если ошибка - выводим сообщение, закрываем соединения, выходим.*/
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
		fmt.Println("Server get message: ", text)
		// Метод Write объекта net.Conn используется для отправки данных клиенту.
		conn.Write([]byte("New server answer: " + text + "\n"))
		if err != nil {
			fmt.Println(err)
			conn.Close()
			return
		}
	}
}
