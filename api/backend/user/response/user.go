package response

type UserLoginResponse struct {
	Token string `json:"token"`
}

type GetUserInfoResponse struct {
	Roles        string `json:"roles"`
	Introduction string `json:"introduction"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
}
