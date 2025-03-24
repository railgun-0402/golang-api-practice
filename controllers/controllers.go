package controllers

import (
	"encoding/json"
	"go-practice-hands/models"
	"go-practice-hands/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MyAppController struct {
	service *services.MyAppService
}

// コンストラクター
func NewMyAppController(s *services.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

// POST /article のハンドラ
func (c *MyAppController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article
	// jsonバイト列をデコード
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	// handler -> service
	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// ArticleListHandler GET /article/list のハンドラ
func (c *MyAppController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

	// handler -> service
	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(articleList)
}

func (c *MyAppController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// Var: パスパラメータをMapで返却する (map[id: 1234])
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])

	// int変換できない場合、つまりパラメータの値が不適切の場合は400エラー
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	// handler -> service
	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice のハンドラ
func (c *MyAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

// POST /comment のハンドラ
func (c *MyAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
