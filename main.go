package main

import (
	"database/sql"
	"errors"
	"fmt"
	"go-practice-hands/api"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser = "user"
	dbPassword = "pass"
	dbPort = "3306"
	dbHost = "localhost"
	dbName = "sampledb"
	dbConn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
)


func main() {

	_, err0 := strconv.Atoi("a")
	fmt.Printf("err0: [%T] %v\n", err0, err0)

	// err0の中に含まれているエラーを取り出し、err1に代入
	err1 := errors.Unwrap(err0)
	fmt.Printf("err1: [%T] %v\n", err1, err1)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}