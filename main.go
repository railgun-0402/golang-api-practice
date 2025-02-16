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

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	const sqlStr = `select title, contents, username, nice from articles;`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArr := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
		} else {
			// 読み出し結果を格納した変数 article を、配列に追加
			articleArr = append(articleArr, article)
		}
	}
	fmt.Printf("%+v\n", articleArr)
}