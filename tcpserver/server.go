package tcpserver

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "8081"
	SERVER_TYPE = "tcp"
)

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
func handleConn(connection net.Conn) {
	buffer := make([]byte, 2048)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	msg := string(buffer[:mLen-1])
	isCloseSignal := msg == ".w_x-q"
	if isCloseSignal {
		fmt.Println("close signal received... closing the server...")
		connection.Write([]byte("We got your close signal... closing server"))
		connection.Close()
		os.Exit(1)
	}
	fmt.Printf("Received message: %s \n", msg)
	_, err = connection.Write([]byte("Get your message as:" + msg))
	connection.Close()
}
