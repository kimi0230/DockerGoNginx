package models

import (
	"time"
)

type User struct {
	ID         int64     `json:"id" form:"id,omitempty" structs:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT:number;type:bigint"`
	Email      string    `json:"email" form:"email,omitempty" structs:"email,omitempty" gorm:"Column:email;type:varchar(30);not null;comment:'信箱'" `
	Password   string    `json:"password" form:"password,omitempty" structs:"password,omitempty" gorm:"Column:password;type:varchar(15);not null;comment:'密碼'" `
	Name       string    `json:"name" form:"name,omitempty" structs:"name,omitempty" gorm:"Column:name;type:varchar(10);not null;comment:'姓名'" `
	Sex        string    `json:"sex" form:"sex,omitempty" structs:"sex,omitempty" gorm:"Column:sex;type:varchar(1);comment:'姓別 0:female,1:male'" `
	AvatarPath string    `json:"avatarPath" form:"avatarPath,omitempty" structs:"avatarPath,omitempty" gorm:"Column:avatarPath;type:text;comment:'圖片路徑'" `
	Birthday   string    `json:"birthday" form:"birthday,omitempty" structs:"birthday,omitempty" gorm:"Column:birthday;type:varchar(10);comment:'生日'" `
	Tel        string    `json:"tel" form:"tel,omitempty" structs:"tel,omitempty" gorm:"Column:tel;type:varchar(20);comment:'電話'"`
	Nation     string    `json:"nation" form:"nation,omitempty" structs:"nation,omitempty" gorm:"Column:nation;type:varchar(3);comment:'國別'"`
	Address    string    `json:"address" form:"address,omitempty" structs:"address,omitempty" gorm:"Column:address;type:text;comment:'地址'"`
	IsActive   string    `json:"isActive" form:"isActive,omitempty" structs:"isActive,omitempty" gorm:"Column:isActive;type:varchar(1);default:'1';comment:'0:關閉, 1:啟用'"`
	ThirdID    string    `json:"third_id" form:"third_id,omitempty" structs:"third_id,omitempty" gorm:"Column:third_id;type:varchar(50);comment:'第三方登入ID'"`
	IsCompany  string    `json:"isCompany" form:"isCompany,omitempty" structs:"isCompany,omitempty" gorm:"Column:isCompany;type:varchar(1);default:'0';comment:'0:非公司,1:公司行號'"`
	CreatedAt  time.Time `gorm:"Column:createdAt"`
	UpdatedAt  time.Time `gorm:"Column:updatedAt"`
}

func (User) TableName() string {
	return "user"
}

// 不讓 ORM 自動把 table name 後面多's'
// func init() {
// 	db.GormDB.SingularTable(true)
// }
