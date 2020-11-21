package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	resp := []byte(`{"name":"Kaioo"}`)
	w.Write(resp)
}

func main() {
	fmt.Println("Start...")
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":1234", nil)

	log.Fatal(err)
	fmt.Println("End.....")
}
