package apputil

const (
	//ResultCodeSuccess API成功
	ResultCodeSuccess = 0

	//ResultCodeRedisError redis実行エラー
	ResultCodeRedisError = 10

	//ResultCodeSessionError セッション切れ
	ResultCodeSessionError = 101
	//ResultCodeCantGetIDFromSession セッションからid取得失敗
	ResultCodeCantGetIDFromSession = 102
	//ResultCodeBANUser セッションからid取得失敗
	ResultCodeBANUser = 103

	//ResultCodeError 予期せぬAPIエラー(なんやそれ)
	ResultCodeError = 500

	//ResultCodeMantenance 全体予定メンテナンス
	ResultCodeMantenance = 9000
	//ResultCodeMantenanceEmergency 全体緊急メンテナンス
	ResultCodeMantenanceEmergency = 9001
)
