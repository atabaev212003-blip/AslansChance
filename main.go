package main

import (
	httpserver "Study/http_server"
	"fmt"
)

func main() {
	fmt.Println("The start of the http server")
	err := httpserver.StartHttpServer()
	if err != nil {
		fmt.Println("Server fell with error :", err)
	} else {
		fmt.Println("Server closed with success!")
	}
}
