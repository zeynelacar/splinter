package main

import (
	"fmt"
	"net/http"
)

const (
	HOST        = "127.0.0.1"
	PORT        = "8088"
	SERVER_TYPE = "http"
)

func main() {
	fmt.Println("Http server running... ")
	http.ListenAndServe(HOST+":"+PORT, nil)
	fmt.Printf("Listening on %s : %s ", HOST, PORT)

}
