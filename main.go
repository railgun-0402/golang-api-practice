package main

import (
	"database/sql"
	"fmt"
	"go-practice-hands/models"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "user"
	dbPassword = "pass"
	dbPort = "3306"
	dbHost = "localhost"
	dbName = "sampledb"
)

func main() {

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	insertArticle := models.Article {
		Title: "insert test",
		Contents: "Can I insert data correctly?",
		UserName: "seki",
	}

	const insertSqlStr = `insert into articles (title, contents, username, nice, created_at) values(?, ?, ?, 0, now());`
	result, err := db.Exec(insertSqlStr, insertArticle.Title, insertArticle.Contents, insertArticle.UserName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result.LastInsertId()) // 何番目のデータかを見る
	fmt.Println(result.RowsAffected()) // DBに何行の変更が行われたか見る

	articleId := 1
	const sqlStr = `select * from articles where article_id = ?;`
	row := db.QueryRow(sqlStr, articleId)

	// データ取得件数が0件の場合は終了
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return
	}

	var article models.Article
	var createdTime sql.NullTime
	err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		fmt.Println(err)
		return
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}
	fmt.Printf("%+v\n", article)
}