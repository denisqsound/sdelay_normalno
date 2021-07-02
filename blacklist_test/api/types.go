package api

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
	Token  string
	Cookie string
}

type AddBlacklistReq struct {
	Bulk []Blacklist `json:"bulk"`
}

type Blacklist struct {
	ClientID int     `json:"clientid"`
	PoolID   int     `json:"poolid"`
	ExpireAt int     `json:"expire_at"`
	IP       string  `json:"ip"`
	Reason   *string `json:"reason"`
}

type AddBlacklistResp struct {
	Body AddBlacklistBody `json:"body"`
}

type AddBlacklistBody struct {
	Result  string      `json:"result"`
	Utils   []int64     `json:"utils"`
	Objects []Blacklist `json:"objects"`
}
