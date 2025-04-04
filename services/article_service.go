package services

import (
	"go-practice-hands/models"
	"go-practice-hands/repositories"
)

// 記事の詳細を取得するサービスクラス
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {


	// repositories層から記事詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	// 次はコメント一覧
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)
	return article, nil
}

// PostArticleHandler で使うことを想定したサービス
// 引数の情報をもとに新しい記事を作り、結果を返却
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	// repositories層から記事データを取得
	articleData, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		return models.Article{}, err
	}

	return articleData, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	article, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		return []models.Article{}, err
	}
	return article, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	dbErr := repositories.UpdateNiceNum(s.db, article.ID)
	if dbErr != nil {
		return models.Article{}, dbErr
	}
	return models.Article{}, nil
}