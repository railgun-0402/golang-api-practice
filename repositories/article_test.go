package repositories_test

import (
	"database/sql"
	"fmt"
	"testing"

	"go-practice-hands/models"
	"go-practice-hands/repositories"

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
// TODO: テストケースを増やす
func TestSelectArticleDetail(t *testing.T) {
	// DB接続
	db, err := connectDB()
	if err != nil {
		// DB接続が失敗したらテスト失敗・終了
		t.Fatal(err)
	}
	defer db.Close()


	expected := models.Article {
		ID: 1,
		Title: "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum: 3,
	}

	// テスト対象クラスを実行
	got, err := repositories.SelectArticleDetail(db, expected.ID)
	if err != nil {
		// 関数の実行自体が失敗：テスト失敗
		t.Fatal(err)
	}

	if got.ID != expected.ID {
		t.Errorf("ID: get %d but want %d\n", got.ID, expected.ID)
	}
	if got.Title != expected.Title {
		t.Errorf("Title: get %s but want %s\n", got.Title, expected.Title)
	}
	if got.Contents != expected.Contents {
		t.Errorf("Content: get %s but want %s\n", got.Contents, expected.Contents)
	}
	if got.UserName != expected.UserName {
		t.Errorf("UserName: get %s but want %s\n", got.UserName, expected.UserName)
	}
	if got.NiceNum != expected.NiceNum {
		t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, expected.NiceNum)
	}

	// 成功

}