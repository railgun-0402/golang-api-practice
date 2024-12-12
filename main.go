package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// HelloWorldを書き込む
		io.WriteString(w, "Hello, World\n")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
