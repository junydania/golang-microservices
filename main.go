package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)
func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops"))
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Println("Goodbye World")
	})
	http.ListenAndServe("127.0.0.1:9090", nil)
}