package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"os"
	"github.com/junydania/golang-microservices/handlers"
)
func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello()

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe("127.0.0.1:9090", sm)
}