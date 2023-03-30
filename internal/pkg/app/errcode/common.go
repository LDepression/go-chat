package errcode

var (
	StatusOK           = NewErr(0, "成功")
	ErrParamsNotValid  = NewErr(1001, "参数有误")
	ErrNotFound        = NewErr(1002, "未找到资源")
	ErrServer          = NewErr(1003, "系统错误")
	ErrTooManyRequests = NewErr(1004, "请求过多")
	ErrTimeOut         = NewErr(1005, "请求超时，请稍后再试")
)

var (
	ErrUnauthorizedAuthNotExist  = NewErr(2001, "鉴权失败,无法解析")
	ErrUnauthorizedTokenTimeout  = NewErr(2002, "鉴权失败,Token超时")
	ErrUnauthorizedTokenGenerate = NewErr(2003, "鉴权失败,Token 生成失败")
	ErrInsufficientPermissions   = NewErr(2004, "鉴权失败,权限不足")
	ErrOutTimeRefreshToken       = NewErr(2005, "refreshToken过期")
	ErrGenerateToken             = NewErr(2006, "生成token失败")
)
