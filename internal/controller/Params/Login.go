package Params

type LoginRec struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Username string `json:"username"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"isAdmin"`
}
