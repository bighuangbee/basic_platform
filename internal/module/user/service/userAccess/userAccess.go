/**
 * @desc //TODO $
 * @param $
 * @return $
 **/
package userAccess

import (
	"errors"
	"github.com/bighuangbee/gokit/tools/cache"
	"time"
)

func New(store cache.Cache, userType string, loginExpire time.Duration) *UserAccess {
	return &UserAccess{Store: store, AccessToken: NewAccessToken(userType), LoginExpire: loginExpire}
}

var Instance *UserAccess

// implement IUserAccess interface
type UserAccess struct {
	Store       cache.Cache
	AccessToken IAccessToken
	LoginExpire time.Duration //登陆有效期
}

func (userService *UserAccess)Issue(user *UserClaims)(string, error){
	claims := user
	claims.ExpiresAt = time.Now().Add(time.Minute * userService.LoginExpire).Unix()

	token, err := userService.AccessToken.Generate(claims)
	if err != nil{
		return "", err
	}
	claims.Token = token

	err = userService.Store.SetEntity(userService.AccessToken.CreateTokenKey(claims), claims, userService.LoginExpire)
	return token, err
}


func (userService *UserAccess) Validate(token string) (*UserClaims, error){
	if token == ""{
		return nil, errors.New("token not allow empty.")
	}

	user, err := userService.AccessToken.Decode(token)
	if err != nil{
		return nil, err
	}

	var validUser UserClaims
	err = userService.Store.GetEntity(userService.AccessToken.CreateTokenKey(user), &validUser)
	if err == nil && validUser.Token == token{
		return &validUser, nil
	}

	return nil, errors.New("token invalid")
}

func (userService *UserAccess)Logout(token string)error{
	user, err := userService.AccessToken.Decode(token)
	if err != nil{
		return err
	}
	return userService.Store.Del(userService.AccessToken.CreateTokenKey(user))
}
