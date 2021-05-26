package config

// key = errCode, string = errMsg
type ErrInfo struct {
	ErrCode int32
	ErrMsg  string
}

var (
	OK = ErrInfo{0, ""}

	ErrMysql             = ErrInfo{100, ""}
	ErrMongo             = ErrInfo{110, ""}
	ErrRedis             = ErrInfo{120, ""}
	ErrParseToken        = ErrInfo{200, "Parse token failed"}
	ErrCreateToken       = ErrInfo{201, "Create token failed"}
	ErrAppServerKey      = ErrInfo{300, "key error"}
	ErrTencentCredential = ErrInfo{400, ""}

	ErrorUserRegister      = ErrInfo{600, "User registration failed"}
	ErrAccountExists       = ErrInfo{601, "The account is already registered and cannot be registered again"}
	ErrUserPassword        = ErrInfo{602, "User password error"}
	ErrTokenIncorrect      = ErrInfo{603, "Invalid token"}
	ErrTokenExpired        = ErrInfo{604, "Expired token"}
	ErrRefreshToken        = ErrInfo{605, "Failed to refresh token"}
	ErrAddFriend           = ErrInfo{606, "Failed to add friends"}
	ErrAgreeToAddFriend    = ErrInfo{607, "Failed to agree to add friend"}
	ErrAddFriendToBlack    = ErrInfo{608, "Failed to add friends to the blacklist"}
	ErrGetBlackList        = ErrInfo{609, "Failed to get blacklist"}
	ErrDeleteFriend        = ErrInfo{610, "Failed to delete friend"}
	ErrGetFriendApplyList  = ErrInfo{611, "Failed to get friend application list"}
	ErrGetFriendList       = ErrInfo{612, "Failed to get friend list"}
	ErrRemoveBlackList     = ErrInfo{613, "Failed to remove blacklist"}
	ErrSearchUserInfo      = ErrInfo{614, "Failed to find user information"}
	ErrDelAppleDeviceToken = ErrInfo{615, ""}
	ErrModifyUserInfo      = ErrInfo{616, "update user some attribute failed"}
	ErrSetFriendComment    = ErrInfo{617, "set friend comment failed"}
)
