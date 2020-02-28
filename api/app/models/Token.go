package models

import (
	config "DockerGoNginx/api/config"
	db "DockerGoNginx/api/config/databases/mysql"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Token struct {
	ID          int64     `json:"id" form:"id,omitempty" structs:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT:number;type:bigint"`
	Token       string    `json:"token" form:"token,omitempty" structs:"token,omitempty" gorm:"Column:token;type:varchar(40);not null;comment:'token' "`
	Securitykey string    `json:"security_key" form:"security_key,omitempty" structs:"security_key,omitempty" gorm:"Column:security_key;type:varchar(40);not null;comment:'安全密碼' "`
	Effect      time.Time `json:"effect" form:"effect,omitempty" structs:"effect,omitempty" gorm:"Column:effect;type:timestamp;default:CURRENT_TIMESTAMP;comment:'有效日' "`
	Expired     time.Time `json:"expired" form:"expired,omitempty" structs:"expired,omitempty" gorm:"Column:expired;type:timestamp;default:CURRENT_TIMESTAMP;comment:'過期日' "`
	UserID      int64     `json:"user_id" form:"user_id,omitempty" structs:"user_id,omitempty" gorm:"Column:user_id;type:bigint;not null;comment:'使用者ID'"`
	CreatedAt   time.Time `gorm:"Column:createdAt"`
	UpdatedAt   time.Time `gorm:"Column:updatedAt"`
}

/*
TableName : 自訂對應的table name
*/
func (Token) TableName() string {
	return "token"
}

/*
BeforeCreate : create時自動建立token跟時間
*/
func (token *Token) BeforeCreate(scope *gorm.Scope) error {
	t := time.Now().UTC()
	expired := t.AddDate(config.TOKEN_TIME[0], config.TOKEN_TIME[1], config.TOKEN_TIME[2])
	scope.SetColumn("Token", uuid.Must(uuid.NewV1()).String())
	scope.SetColumn("Securitykey", uuid.Must(uuid.NewV4()).String())
	scope.SetColumn("Effect", t)
	scope.SetColumn("Expired", expired)
	return nil
}

/*
GetUser : 依據token取得 user model
*/
func (token *Token) GetUser() User {
	var userModel User
	db.GormDB.Where("id = ?", token.UserID).First(&userModel)
	return userModel
}
