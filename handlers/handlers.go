package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// ArticleListHandler GET /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// reqから、クエリパラメータの値を取り出したい
	queryMap := req.URL.Query()

	var page int
	// パラメータ page が1個以上あるなら
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	resString := fmt.Sprintf("ArticleList (page %d)\n", page)
	io.WriteString(w, resString)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// Var: パスパラメータをMapで返却する (map[id: 1234])
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	// int変換できない場合、つまりパラメータの値が不適切の場合は400エラー
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func ArticleNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

func ArticleCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
