// Code generated by goctl. DO NOT EDIT.
package types

type CrudReq struct {
	Object string `form:"object"`
	Action string `form:"action"`
	Data   string `form:"data"`
}

type CrudRes struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
	Ok   bool   `json:"ok"`
}
