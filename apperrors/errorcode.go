package apperrors

type ErrCode string

const (
	// とりあえず想定外のエラーコードを投入！
	Unknown ErrCode = "U000"

	// Insert失敗
	InsertDataFailed ErrCode = "S001"
	// Select失敗(処理失敗)
	GetDataFailed ErrCode = "S002"
	// Select失敗(データ0件)
	NAData ErrCode = "S003"
	// 該当データなし
	NoTargetData ErrCode = "S004"
	// 更新失敗
	UpdateDataFailed ErrCode = "S005"

	// Controller層で出るエラー
	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam ErrCode = "R002"
)