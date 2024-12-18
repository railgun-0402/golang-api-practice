package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		// HelloWorldを書き込む
		io.WriteString(w, "Hello, World\n")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Artcile...\n")
	}

	articleListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Article List\n")
	}

	articleDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("Article No.%\n", articleID)
		io.WriteString(w, resString)
	}

	articleNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Nice...\n")
	}

	articleCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Posting Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", articleListHandler)
	http.HandleFunc("/article/1", articleDetailHandler)
	http.HandleFunc("/article/nice", articleNiceHandler)
	http.HandleFunc("/comment", articleCommentHandler)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
