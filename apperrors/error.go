package apperrors

type MyAppError struct {
	// 独自エラーに含めるフィールドの定義
	ErrCode // レスポンスとログに表示するエラーコード
	Message string // レスポンスに表示するエラーメッセージ
	Err error  // エラーチェーンのための内部エラー
}

func (myErr *MyAppError) Error() string {
	return myErr.Err.Error()
}

// 入れ子にして含む内部エラーを返す
func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}

func (code ErrCode) Wrap(err error, message string) error {
	return &MyAppError{ErrCode: code, Message: message, Err: err}
}