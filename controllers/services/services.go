package services

import "go-practice-hands/models"

// article関連を引き受ける
type ArticleServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	GetArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)
}

// comment関連を引き受ける
type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
