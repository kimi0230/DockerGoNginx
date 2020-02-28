package models

import "time"

type ApiLog struct {
	Id            int64     `json:"id" form:"id,omitempty" structs:"id,omitempty"`
	UserID        int       `json:"user_id" form:"user_id,omitempty" structs:"user_id,omitempty" gorm:"Column:user_id;type:int;comment:'使用者ID' "`
	RequestMethod string    `json:"request_method" form:"request_method,omitempty" structs:"request_method,omitempty" gorm:"Column:request_method;type:varchar(5);comment:'request method'"`
	ApiURL        string    `json:"api_url" form:"api_url,omitempty" structs:"api_url,omitempty" gorm:"Column:api_url;type:varchar(100);comment:'api url'"`
	ResultCode    string    `json:"error_code" form:"error_code,omitempty" structs:"error_code,omitempty" gorm:"Column:error_code;type:varchar(5);comment:'error code'"`
	ResultMsg     string    `json:"result_msg" form:"result_msg,omitempty" structs:"result_msg,omitempty" gorm:"Column:result_msg;type:text;comment:'error result_msg'"`
	Request       string    `json:"request" form:"request,omitempty" structs:"request,omitempty" gorm:"Column:request;type:text;comment:'api request'"`
	Response      string    `json:"response" form:"response,omitempty" structs:"response,omitempty" gorm:"Column:response;type:text;comment:'api response'"`
	ErrorMsg      string    `json:"error_msg" form:"error_msg,omitempty" structs:"error_msg,omitempty" gorm:"Column:error_msg;type:text;comment:'error error_msg'"`
	IP            string    `json:"ip" form:"ip,omitempty" structs:"ip,omitempty" gorm:"Column:ip;type:text;comment:''"`
	CreatedAt     time.Time `json:"createdAt" form:"createdAt,omitempty" structs:"createdAt,omitempty" gorm:"Column:createdAt;"`
	UpdatedAt     time.Time `json:"updatedAt" form:"updatedAt,omitempty" structs:"updatedAt,omitempty" gorm:"Column:updatedAt;"`
}

// 自訂對應的table name
func (ApiLog) TableName() string {
	return "api_log"
}
