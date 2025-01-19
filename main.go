package main

import (
	"encoding/json"
	"fmt"
	"go-practice-hands/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Comment struct {
	CommentID int        `json:"comment_id"`
	ArticleID int        `json:"article_id"`
	Message   string     `json:"message"`
	CreatedAt time.Time  `json:"created_at"`
}

type Article struct {
	ID          int        `json:"article_id"`
	Title       string     `json:"title"`
	Contents    string     `json:"contents"`
	UserName    string     `json:"user_name"`
	NiceNum     int        `json:"nice"`
	CommentList []Comment  `json:"comments"`
	CreatedAt   time.Time  `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}

	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article",
		UserName:    "su",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	// jsonエンコード
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)

	// Router作成
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.ArticleNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.ArticleCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
