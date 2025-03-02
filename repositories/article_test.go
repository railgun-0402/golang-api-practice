package repositories_test

import (
	"fmt"
	"testing"

	"go-practice-hands/models"
	"go-practice-hands/repositories"
	"go-practice-hands/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected models.Article
	}{
		{
			testTitle: "subtest1",
			expected: testdata.ArticleTestData[0],
		},
		{
			testTitle: "subtest2",
			expected: testdata.ArticleTestData[1],
		},
	}

	for _, test := range tests {
		// Runでサブテストを書く
		t.Run(test.testTitle, func(t *testing.T) {
			// テスト対象クラスを実行
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				// 関数の実行自体が失敗：テスト失敗
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}
	// 成功

}

// SelectArticleList関数のテスト
func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// InsertArticle関数のテスト
func TestInsertArticle(t *testing.T) {
	article := models.Article {
		Title: "insertTest",
		Contents: "testtest",
		UserName: "saki",
	}

	expectedArticleNum := 10
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Error(err)
	}

	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		fmt.Println("delete start")
		const sqlStr = `
		delete from articles
		where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
		fmt.Println("delete end")
	})
}

