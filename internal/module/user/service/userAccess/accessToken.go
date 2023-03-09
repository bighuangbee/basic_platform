/**
 * @desc //TODO $
 * @param $
 * @return $
 **/
package userAccess

import (
	"fmt"
	"github.com/bighuangbee/gokit/tools"
	"github.com/bighuangbee/gokit/tools/crypto"
	"time"
)

type IAccessToken interface {
	//生成token
	Generate(claims *UserClaims)(token string, err error)

	Decode(token string)(claims *UserClaims, err error)

	CreateTokenKey(claims *UserClaims) (string)
}


type BaseToken struct{
	Type 		string	//用户类型
	Encrtpy 	string	//密钥
	TokenKey 	string	//token索引
	InValidTokenKey string	//失效token索引
}


func NewBaseToken(accessType string, encrpty string) *BaseToken {
	return &BaseToken{
		Type:            accessType,
		Encrtpy:         encrpty,
		TokenKey:        accessType + ":token:%s:%d:%d", //$accessType:token:$username:$platform:$userType
		InValidTokenKey: accessType + ":tokenInValid:%s:%d:%d", //$accessType:tokenInValid:$username:$platform:$userType
	}
}

func NewAccessToken(userType string) IAccessToken {
	//return NewAccessTokenJWT(userType)
	return NewAccessTokenCommon(userType)
}

func (this *BaseToken) Generate(claims *UserClaims) (string, error) {
	str := claims.UserName + "#" + fmt.Sprintf("%d", claims.Platform) +"#" +
		fmt.Sprintf("%d", claims.Type) + "#" + tools.MD5(this.Encrtpy + "_HIDRONE_" + time.Now().String())
	return crypto.AesEncryptStr(str, this.Encrtpy)
}

func (this *BaseToken)CreateTokenKey(claims *UserClaims) (string){
	return fmt.Sprintf(this.TokenKey, claims.UserName, claims.Platform, claims.Type)
}
