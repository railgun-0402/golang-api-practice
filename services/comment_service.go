package services

import (
	"go-practice-hands/models"
	"go-practice-hands/repositories"
)

// 引数の情報をもとに新しいコメントを作り、結果を返却
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {

	comments, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return comments, nil
}