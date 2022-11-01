package main

import (
	"fmt"
	"splinter/tcpserver"
)

func init() {
	fmt.Println("initialized")
}

func main() {
	tcpserver.Start()
}
