package main

import (
	"database/sql"
	"fmt"
	"go-practice-hands/controllers"
	"go-practice-hands/routers"
	"go-practice-hands/services"
	"log"
	"net/http"

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
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	service := services.NewMyAppService(db)
	con := controllers.NewMyAppController(service)

	r := routers.NewRouter(con)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}