package services

import (
	"go-practice-hands/models"
	"go-practice-hands/repositories"
)

// PostCommentHandler で使用することを想定したサービス
// 引数の情報をもとに新しいコメントを作り、結果を返却
func PostCommentService(comment models.Comment) (models.Comment, error) {
	db, err := connectDB()
	if err != nil {
		return models.Comment{}, err
	}
	defer db.Close()

	comments, err := repositories.InsertComment(db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return comments, nil
}