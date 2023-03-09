package userAccess

import (
	"errors"
	"fmt"
	"github.com/bighuangbee/gokit/tools/crypto"
	"strconv"
	"strings"
)

const CommonEncrtpy = "4cdc1f67bf13aae1fec30fdb3f721380"

type AccessTokenCommon struct{
	*BaseToken
}

func NewAccessTokenCommon(accessType string) IAccessToken {
	return &AccessTokenCommon{
		NewBaseToken(accessType, CommonEncrtpy),
	}
}

func (this *AccessTokenCommon) Generate(claims *UserClaims) (string, error) {
	return this.BaseToken.Generate(claims)
}

func (this *AccessTokenCommon)Decode(token string)(claims *UserClaims, err error){
	str, err := crypto.AesDecryptStr(token, this.Encrtpy)
	if err != nil{
		return &UserClaims{}, err
	}

	decodeStr := strings.Split(str, "#")
	if len(decodeStr) >= 4{
		platform, _ := strconv.Atoi(decodeStr[1])
		userType, _ := strconv.Atoi(decodeStr[2])
		return &UserClaims{
			UserName:       decodeStr[0],
			Platform:       int8(platform),
			Type:       	int8(userType),
		}, err
	}
	return nil, errors.New("Decode err")
}

func (this *AccessTokenCommon)CreateTokenKey(claims *UserClaims) (string){
	return fmt.Sprintf(this.TokenKey, claims.UserName, claims.Type, claims.Platform)
}
