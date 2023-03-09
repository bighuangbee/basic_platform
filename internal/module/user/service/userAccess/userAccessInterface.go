/**
 * @desc //TODO $
 * @param $
 * @return $
 **/
package userAccess


type IUserAccess interface {
	//签发用户访问凭证 token
	Issue(*UserClaims) (token string, err error)

	//对凭证进行认证
	Validate(token string)(*UserClaims, error)

	//注销访问凭证
	Logout(username string)(error)
}
