package main

import (
	"log"
	"net/http"
)

// Server - http server
var Server *http.Server

func main() {
	Server := newServer()
	log.Printf("Start serving on %s", Server.Addr)
	log.Println(Server.ListenAndServe())
}
