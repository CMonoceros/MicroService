package errcode

var (
	OK = add(0) // 正确

	ServerErr          = add(-500) // 服务器错误
	ServiceUnavailable = add(-503) // 服务暂不可用
)
