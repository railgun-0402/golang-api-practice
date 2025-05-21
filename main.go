package main

import (
	"database/sql"
	"fmt"
	"go-practice-hands/api"
	"io"
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

	r := api.NewRouter(db)

	helloHandler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	http.Handle("/", myMiddleware1(helloHandler))
	log.Fatal(http.ListenAndServe(":8080", r))
}

func myMiddleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Pre-process1\n")
		next.ServeHTTP(w, r)
		io.WriteString(w, "Post-process1\n")
	})
}