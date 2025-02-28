package repositories_test

import (
	"testing"

	"go-practice-hands/models"
	"go-practice-hands/repositories"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected models.Article
	}{
		{
			testTitle: "subtest1",
			expected: models.Article {
				ID: 1,
				Title: "firstPost",
				Contents: "This is my first blog",
				UserName: "saki",
				NiceNum: 3,
			},
		},
		{
			testTitle: "subtest2",
			expected: models.Article{
				ID: 2,
				Title: "2nd",
				Contents: "Second blog post",
				UserName: "saki",
				NiceNum: 4,
			},
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
	expectedNum := 3
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

