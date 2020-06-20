package main

import (
	"net/http"
	"log"
	"os"
	"github.com/junydania/golang-microservices/handlers"
)
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe("127.0.0.1:9090", sm)
}