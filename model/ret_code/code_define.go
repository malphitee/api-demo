package ret_code

var (
	PARAMS_FORMAT_INVALID = ErrorResponse{-2000, "请求格式不正确", "params format error"}
	PARAMS_INVALID        = ErrorResponse{-2001, "请求参数不合法", "params invalid"}
	PERMISSION_DENY       = ErrorResponse{-2002, "没有权限", "permission deny"}
	SIGNATURE_ERROR       = ErrorResponse{-2003, "签名错误", "signature error"}
	SYSTEM_ERROR          = ErrorResponse{-2004, "系统错误", "system error"}
)
