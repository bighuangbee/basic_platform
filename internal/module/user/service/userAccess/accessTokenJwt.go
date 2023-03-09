package userAccess

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const JWT_Encrtpy = "7233kweGJdAAzy13daa6"

type AccessTokenJWT struct{
	*BaseToken
}

func NewAccessTokenJWT(userType string) IAccessToken {
	return &AccessTokenJWT{
		NewBaseToken(userType, JWT_Encrtpy),
	}
}

func (this *AccessTokenJWT) Generate(claims *UserClaims) (string, error) {
	return this.BaseToken.Generate(claims)
}

func (this *AccessTokenJWT)Decode(tokenStr string)(claims *UserClaims, err error){
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(this.Encrtpy), nil
	})

	if err == nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func (this *AccessTokenJWT)CreateTokenKey(claims *UserClaims) (string){
	return this.BaseToken.CreateTokenKey(claims)
}
