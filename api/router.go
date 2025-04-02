package api

import (
	"database/sql"
	"go-practice-hands/controllers"
	"go-practice-hands/services"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {

	ser := services.NewMyAppService(db)
	articleConn := controllers.NewArticleController(ser)
	commentConn := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/article", articleConn.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", articleConn.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", articleConn.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", articleConn.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", commentConn.PostCommentHandler).Methods(http.MethodPost)
	return r
}