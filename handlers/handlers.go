package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// GETメソッドなら
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, World!\n")
	} else {
		// GETメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		// POSTメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		// GETメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	if req.Method == http.MethodGet {
		io.WriteString(w, resString)
	} else {
		// GETメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		// POSTメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func ArticleCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		// POSTメソッド以外は405
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
