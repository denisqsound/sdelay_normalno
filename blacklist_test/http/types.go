package http

const (
	BaseUrl         = "http://172.25.84.30:9752"
	LoginUrl        = "%s/v1/login"
	AddBlacklistUrl = "%s/v3/blacklist/bulk"
)

type AuthResp struct {
	StatusCode int  `json:"status"`
	Body       Body `json:"body"`
}

type Body struct {
	Token string `json:"token"`
}

type AuthReg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthorisedUser struct {
	token  string
	cookie string
}
