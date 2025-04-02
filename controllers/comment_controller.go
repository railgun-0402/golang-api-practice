package controllers

import (
	"encoding/json"
	"go-practice-hands/controllers/services"
	"go-practice-hands/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

// コンストラクター
func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
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

