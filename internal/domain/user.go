package domain

import (
	"context"
	"github.com/bighuangbee/gokit/model"
)

type User struct {
	ID        	model.Id
	Username	string `gorm:"column:username" json:"username"`
	Nickname	string `gorm:"column:nickname" json:"nickname"`
	Mobile       string `gorm:"column:mobile" json:"mobile"`
	Email       string `gorm:"column:mobile" json:"email"`
	HeadPicUrl  string `gorm:"column:head_pic_url" json:"head_pic_url"`
	Password	string `gorm:"column:password" json:"password,omitempty"`
	Salt		string `gorm:"column:salt" json:"-"`
	Status		int8 `gorm:"column:status" json:"status"`
	Ip			string `gorm:"column:ip" json:"ip"`
	model.CreatedInfo
	model.UpdatedInfo
	model.DeletedInfo
}

type UserLogin struct {
	Username	string `gorm:"column:username" json:"username"`
	Password	string `gorm:"column:password" json:"password,omitempty"`
	Mobile      string `gorm:"column:mobile" json:"mobile"`
}

type IUserRepo interface {
	Login(context.Context, *User) (*User, error)
	CreateUser(context.Context, *User) (*User, error)
}
