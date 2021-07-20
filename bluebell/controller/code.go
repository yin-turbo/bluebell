package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
)

//定义Msg map
var codeMsg = map[ResCode]string{}
