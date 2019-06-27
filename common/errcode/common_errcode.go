package errcode

var (
	OK = add(0, "OK")

	ServerErr          = add(-500, "服务器错误")
	ServiceUnavailable = add(-503, "服务暂不可用")

	OssConfigError   = add(-1001, "OSS配置错误")
	OssResourceError = add(-1002, "OSS资源错误")
)
