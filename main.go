package main

import (
	"go-practice-hands/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.ArticleListHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	http.HandleFunc("/article/nice", handlers.ArticleNiceHandler)
	http.HandleFunc("/comment", handlers.ArticleCommentHandler)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
