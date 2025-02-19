package main

import (
	"database/sql"
	"fmt"
	"go-practice-hands/models"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "user"
	dbPassword = "pass"
	dbPort = "3306"
	dbHost = "localhost"
	dbName = "sampledb"
)

// DB接続
func connectDB() (*sql.DB, error) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, fmt.Errorf("DB接続エラー: %w", err)
	}
	return db, nil
}

// 記事の挿入
func insertArticleFunc(db *sql.DB, article models.Article) (sql.Result, error) {
	const insertSqlStr = `insert into articles (title, contents, username, nice, created_at) values(?, ?, ?, 0, now());`
	result, err := db.Exec(insertSqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return nil, fmt.Errorf("データのINSERTが失敗しました: %w", err)
	}
	return result, nil
}

// IDでDBからデータを取得する
func getArticleByID(db *sql.DB, articleId int) (*models.Article, error) {
	const sqlStr = `select * from articles where article_id = ?;`
	row := db.QueryRow(sqlStr, articleId)

	// データ取得件数が0件の場合は終了
	if err := row.Err(); err != nil {
		return nil, fmt.Errorf("データが見つかりませんでした: %w", err)
	}

	var article models.Article
	var createdTime sql.NullTime
	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return nil, fmt.Errorf("データの取得に失敗しました: %w", err)
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return &article, nil
}

func main() {

	// DBに接続
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 追加するデータ
	// insertArticle := models.Article {
	// 	Title: "insert test",
	// 	Contents: "Can I insert data correctly?",
	// 	UserName: "seki",
	// }

	// データを追加する
	// result, err := insertArticleFunc(db, insertArticle)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result.LastInsertId()) // 何番目のデータかを見る
	// fmt.Println(result.RowsAffected()) // DBに何行の変更が行われたか見る

	articleId := 1
	// IDに紐づくデータを取得する
	article, err := getArticleByID(db, articleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", article)

	// データのUpdate
	// TODO: 関数化

	// トランザクション開始
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 現在のいいね数を取得するクエリを実行する
	article_id := 1
	const sqlGetNice = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlGetNice, article_id)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// クエリの結果を変数に突っ込む
	var niceNum int
	err = row.Scan(&niceNum)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	// いいね数を+1
	const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
	_, err = tx.Exec(sqlUpdateNice, niceNum + 1, article_id)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return
	}

	tx.Commit()
}