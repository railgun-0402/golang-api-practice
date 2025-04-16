package services

import (
	"database/sql"
	"errors"
	"go-practice-hands/apperrors"
	"go-practice-hands/models"
	"go-practice-hands/repositories"
)

// 記事の詳細を取得するサービスクラス
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {

	// repositories層から記事詳細を取得
	article, err := repositories.SelectArticleDetail(s.db, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Article{}, err
	}

	// 次はコメント一覧
	commentList, err := repositories.SelectCommentList(s.db, articleID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
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
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Article{}, err
	}

	return articleData, nil
}

// ArticleListHandler で使うことを想定したサービス
// 指定 page の記事一覧を返却
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {

	article, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return []models.Article{}, err
	}

	if len(article) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return article, nil
}

// PostNiceHandler で使うことを想定したサービス
// 指定 ID の記事のいいね数を+1 して、結果を返却
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {

	dbErr := repositories.UpdateNiceNum(s.db, article.ID)
	if dbErr != nil {
		if errors.Is(dbErr, sql.ErrNoRows) {
			dbErr = apperrors.NoTargetData.Wrap(dbErr, "does not exist target article")
			return models.Article{}, dbErr
		}
		dbErr = apperrors.UpdateDataFailed.Wrap(dbErr, "fail to update nice count")
		return models.Article{}, dbErr
	}
	return models.Article{}, nil
}