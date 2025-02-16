package main

import (
	"database/sql"
	"fmt"
	"go-practice-hands/models"
	"time"

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

	const sqlStr = `select * from articles;`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArr := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)

		// Nullでなければ、time.Timeに変換
		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		} else {
			article.CreatedAt = time.Time{}
		}

		if err != nil {
			fmt.Println(err)
		} else {
			// article を articleArr に追加
			articleArr = append(articleArr, article)
		}
	}
	fmt.Printf("%+v\n", articleArr)
}