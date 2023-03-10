package domain

import (
	"context"
	v1 "github.com/bighuangbee/basic-platform/api/user/v1"
	"github.com/bighuangbee/gokit/model"
)

type IUserRepo interface {
	Login(context.Context, *User) (*User, error)
	Create(context.Context, *User) (error)
	Update(context.Context, *v1.UpdateUserRequest) (error)
}

type User struct {
	//model.Id
	Id int64 `gorm:"type:bigint(11) not null auto_increment;primaryKey" json:"id"`
	Account string `gorm:"column:account" json:"account"`
	UserName	string `gorm:"column:user_name" json:"user_name"`
	Mobile      *string `gorm:"column:mobile" json:"mobile"`
	Email       *string `gorm:"column:email" json:"email"`
	Password	string `gorm:"column:password" json:"password,omitempty"`
	Salt		string `gorm:"column:salt" json:"-"`
	Status		int8 `gorm:"column:status" json:"status"`
	model.CreatedInfo
	model.UpdatedInfo
	model.DeletedInfo
	DeletedUnix	int64 `json:"-"`

	//Profile UserProfile `gorm:"foreignkey:id;references:user_id"`
}

type UserLogin struct {
	Account		string `json:"account"`
	Password	string `json:"password,omitempty"`
}
