package util

func VerifyToken(token string, scope string) (isAuth bool, plainUserInfo string) {

	encryptedUserInfo := QueryUserInfoByTokenAndScope(token, scope)
	if encryptedUserInfo == "" {
		return false, ""
	}
	plainUserInfo = Decrypt(encryptedUserInfo)

	return true, plainUserInfo
}
