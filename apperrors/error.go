package apperrors

type MyAppError struct {
	// 独自エラーに含めるフィールドの定義
	ErrCode // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
}

func (myErr *MyAppError) Error() string {
	return myErr.Message
}