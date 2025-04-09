package apperrors

type ErrCode string

const (
	// とりあえず想定外のエラーコードを投入！
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed ErrCode = "S002"
	NAData ErrCode = "S003"
)