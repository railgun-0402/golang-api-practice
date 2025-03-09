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
func PostArticleService(article models.Article) (models.Article, error) {
	// TODO : 実装
	return models.Article{}, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func GetArticleListService(page int) ([]models.Article, error) {
	// TODO : 実装
	return nil, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func PostNiceService(article models.Article) (models.Article, error) {
	// TODO : 実装
	return models.Article{}, nil
}