package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func todosHandler(w http.ResponseWriter, req *http.Request) {

	switch method := req.Method; method {
	case "GET":
		fmt.Fprintf(w, "hello GET todos\n")
	case "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "Err = %v\n", err)
			return
		}

		//Set body in response to json
		w.Header().Set("Content-Type", "application/json")

		data := []byte(body)
		t := Todo{}
		json.Unmarshal(data, &t)

		fmt.Fprintf(w, "%s", body)
		fmt.Printf("Body = %#v\n", t)

	case "DELETE":
		fmt.Fprintf(w, "hello DELETE todos\n")
	case "PUT":
		fmt.Fprintf(w, "hello PUT todos\n")
	default:
		fmt.Fprintf(w, "hello")
	}
}

func main() {
	fmt.Println("Start...")
	http.HandleFunc("/todos", todosHandler)
	log.Fatal(http.ListenAndServe(":1234", nil))

	fmt.Println("End.....")
}
