package errorCode

/* 錯誤代碼回傳*/

var PARAMS_INVALID = map[string]string{
	"code":    "1003",
	"message": "必填、不可填的參數有誤",
}
var SUCCESS = map[string]string{
	"code":    "2000",
	"message": "正確狀況",
}
var TOKEN_INVALID = map[string]string{
	"code":    "3001",
	"message": "沒有正確給予Token",
}
var TOKEN_EXPIRED = map[string]string{
	"code":    "3009",
	"message": "Token逾期 不可使用",
}
var ID_ABNORMAL = map[string]string{
	"code":    "5001",
	"message": "使用者身分異常",
}
var USER_INVALID = map[string]string{
	"code":    "5002",
	"message": "使用者停用 或 未開通",
}
var THIRDID_ABNORMAL = map[string]string{
	"code":    "5007",
	"message": "第三方登入id不正確",
}
var DB_WRITE_FAIL = map[string]string{
	"code":    "9001",
	"message": "寫入DB資料錯誤回報",
}
var DB_UPDATE_FAIL = map[string]string{
	"code":    "9003",
	"message": "修改DB資料錯誤回報",
}
var NATION_INVALID = map[string]string{
	"code":    "9009",
	"message": "國別為空或填寫錯誤",
}

/*
參數問題	1001	RESTful 的 DirectType 不正確
參數問題	1002	預設參數，輸入格式不正確
參數問題	1003	必填、不可填的參數有誤
參數問題	1004	轉換欄位參數缺乏
參數問題	1005	第三方的公司名稱，輸入錯誤
加好友	1101	使用者不存在
加好友	1102	已加過好友
加好友	1103	正在邀請中,等待確認
加好友	1104	只能刪除自己邀請的朋友
加好友	1105	不可加自己好友
加好友	1106	無邀請紀錄
正確情形	2000	正確狀況
正確情形	2001	正確狀況，但查無資料
正確情形	2002	SQL查詢正常
正確情形	2003	預定查詢
正確情形	2004	Device fw已經是最新版
Token問題	3001	沒有正確給予Token
Token問題	3002	Token查詢DB異常
Token問題	3003	Token查無對應人員資料
Token問題	3004	token對應人員，權限不符合
Token問題	3005	帳密找Token給予的參數不足
Token問題	3006	帳密尋找TOKEN，錯誤回報
Token問題	3007	帳密查無資料
Token問題	3008	帳密來源有問題
Token問題	3009	Token逾期 不可使用
Token問題	3010	token對應人員，創建權限不符
AWS問題	4001	上傳S3，錯誤回報
身分問題	5001	使用者身分異常
身分問題	5002	使用者停用 或 未開通
身分問題	5003	後台使用者 停用 或 未開通
身分問題	5004	使用者已存在，不可創建。
身分問題	5005	使用者沒有與裝置綁定或者到期
身分問題	5006	使用者不是owner不可控制此裝置
身分問題	5007	第三方登入id不正確
身分問題	5008	請先建立密碼
身分問題	5009	mac_address不正確
身分問題	5010	device owner超過1個人
身分問題	5011	第三方登入email不正確
身分問題	5012	被分享裝置的使用者停用或未開通
挖寶問題	6001	此地的寶箱已經被人捷足先登了喔。
裝置問題	7001	Device異常 可能是狀態問題
裝置問題	7002	Device與Sensor 綁定異常
參數問題	7003	沒填寫語言 lang
參數問題	7004	沒填寫時區
裝置問題	7005	Device version 格式錯誤
裝置問題	7006	Device 不是gateway
裝置問題	7007	Device 不允許分享
裝置問題	7008	配信已滿-不發Access KEY
裝置問題	7009	請進行藍芽reset
裝置問題	7010	Device已被綁定
裝置問題	7011	device_binding不可刪除自己
裝置問題	7012	裝置已經分享中,等待對方同意
裝置問題	7013	選擇的鈴聲已重複
裝置問題	7014	Device驅動驅動更新失敗 (原7004改7014)
裝置問題	7015	裝置未連線
裝置問題	7016	裝置沒有綁定或到期
裝置問題	7017	配信異常
MQTT	8001	MQTT Server Timeout
MQTT	8002	MQTT push 失敗
MQTT	8003	MQTT 帳密錯誤
MQTT	8004	client id異常
MQTT	8005	MQTT collar did 錯誤
MQTT	8006	MQTT_pw錯誤,請要求更換新密碼
錯誤回傳	9001	寫入DB資料錯誤回報
錯誤回傳	9002	查詢DB資料錯誤回報
錯誤回傳	9003	修改DB資料錯誤回報
錯誤回傳	9004	停用DB資料錯誤回報
錯誤回傳	9005	SQL error
錯誤回傳	9006	特殊異常錯誤
錯誤回傳	9007	紀錄失敗 (有可能是紀錄DB crash)
錯誤回傳	9008	客製API錯誤
錯誤回傳	9009	國別為空或填寫錯誤
錯誤回傳	9010	gcm發送失敗
特殊錯誤回傳	8888	( 特殊文字回應 撈取data) + 跳回首頁
特殊錯誤回傳	9999	( 特殊文字回應 撈取data)
*/
