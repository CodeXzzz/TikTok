package model

// ResMsg 响应json信息的结构体
type ResMsg struct {
	Code string      `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
