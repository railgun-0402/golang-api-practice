package services

import (
	"go-practice-hands/models"
	"go-practice-hands/repositories"
)

// 記事の詳細を取得するサービスクラス
func GetArticleService(articleID int) (models.Article, error) {

	db, err := connectDB()
	if err != nil {
		return models.Article{}, err
	}
	defer db.Close()

	// repositories層から記事詳細を取得
	article, err := repositories.SelectArticleDetail(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 次はコメント一覧
	commentList, err := repositories.SelectCommentList(db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// TODO: 記事やいいね数を取得するServiceメソッドを作成する