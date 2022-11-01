package code

type MsgCode uint

const (
	Success MsgCode = iota + 0x0000
	Failed
	Logined
	NoExist
)
