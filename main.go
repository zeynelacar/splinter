package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"unicode/utf8"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "8081"
	SERVER_TYPE = "tcp"
)

func init() {
	fmt.Println("initialized")
}

func main() {
	Start()
}

func Start() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening address:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Printf("Listening on %s : %s ", SERVER_HOST, SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error creating connection: ", err.Error())
			os.Exit(1)
		}
		ipAddr := connection.RemoteAddr().String()
		fmt.Printf("client connected from: %s \n", ipAddr)
		go handleConn(connection)
	}
}

func processMessage(msg string) (res []string) {
	length := msg[:6]
	strL, err := strconv.ParseInt(string(length), 16, 64)
	if err != nil {
		println(err)
		res = append(res, "Unsuccessful")
		res = append(res, "Invalid Payload")
	}
	if int64(utf8.RuneCountInString(msg[6:])) != strL {
		println(utf8.RuneCountInString(msg[6:]))
		println(strL)
		println("Message payload and length are in mismatch")
		res = append(res, "Unsuccessful")
		res = append(res, "Invalid Payload")
	}
	res = append(res, "Successful")
	res = append(res, msg[6:])
	return res
}

func handleConn(connection net.Conn) {
	buffer := make([]byte, 2048)
	errflag := 0
	msg := ""
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		errflag = 1
	}
	if errflag == 0 {
		msg = string(buffer[:mLen-1])
	}
	processed := processMessage(msg)
	processResult := processed[0]
	payload := processed[1]
	isCloseSignal := payload == ".w_x-q"
	if isCloseSignal {
		fmt.Println("close signal received... closing the server...")
		connection.Write([]byte("We got your close signal... closing server"))
		connection.Close()
		os.Exit(1)
	}
	fmt.Printf("Received message: %s \n", payload)
	dump := []byte("Got your message as:" + payload + " sent result :" + processResult)
	connection.Write(dump)
	connection.Close()
}
