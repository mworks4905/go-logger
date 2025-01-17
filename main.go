package main

import (
	"log"

	"github.com/mworks4905/go-logger/internal/server"
)

func main() {
	svr := server.NewHTTPServer(":8080")
	log.Fatal(svr.ListenAndServe())
}
