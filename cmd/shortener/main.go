package main

import (
	"fmt"

	"github.com/vadim-ivlev/url-shortener/internal/server"
)

func main() {
	fmt.Println("Server starting at http://localhost:8080/")
	server.Serve()
}
