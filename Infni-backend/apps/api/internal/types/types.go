// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.2

package types

type LoginReq struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

type LoginResp struct {
	UserId   int64  `json:"userId"`
	Name     string `json:"name"`
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}

type RegisterReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	VerifyCode string `json:"verifyCode"`
}

type RegisterResp struct {
}
