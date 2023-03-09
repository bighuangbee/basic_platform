package userAccess

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const ROLE_APP = "app"
const ROLE_SUPER_ADMIN = "superadmin"

const USE_ACCESS_SYS = "sysUser"
const USE_ACCESS_APP = "user"
const USE_ACCESS_DRONE = "droneUser"

const USER_TYPE_MOINTOR = 1
const USER_TYPE_DUTY = 2
const USER_TYPE_DRONE = 3

const PlatformAccount = 0
const PlatformWx = 1
const PlatformWxUnion = 2

var USER_TYPE  = map[int]string{
	USER_TYPE_MOINTOR: USE_ACCESS_SYS,
	USER_TYPE_DUTY:    USE_ACCESS_APP,
	USER_TYPE_DRONE:   USE_ACCESS_DRONE,
}

type UserClaims struct {
	UserName string   `json:"username"`
	NickName string   `json:"nickname"`
	Type     int8     `json:"type"` //用户类型
	CId      int32    `json:"c_id"` //公司ID
	CName    string    `json:"c_name"` //公司唯一标识
	Roles    []string `json:"roles,omitempty"`
	jwt.StandardClaims
	Token	string `json:"token,omitempty"`
	Platform int8 `json:"platform"`
}

func NewUserClaims(userName string, nickName string, userType int8, companyId int32, roles []string, Id string, expire time.Duration) *UserClaims {
	return &UserClaims{UserName: userName, NickName: nickName, Type: userType, CId: companyId, Roles: roles, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(expire)).Unix(),
		Id:        Id,
	}}
}
