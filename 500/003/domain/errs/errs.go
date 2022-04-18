package errs

import (
	util "ddd/domain/utility"
	"fmt"
)

type ErrorCoder interface {
	Error() ErrorCode
}

// 错误信息
type ErrorCode struct {
	Code    int    `json:"code"`    // 错误码
	Message string `json:"message"` // 错误信息
}

const ErrorCodeStart = 10001

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
const (
	// HTTP 400类
	ErrorParamInvalid = ErrorCodeStart + iota // 无效参数

	// HTTP 404类
	ErrorResourceNotExist // 资源不存在

	ErrorResourceIsExist    // 资源已存在，不允许重复添加
	ErrorCapNotSupport      // 能力不支持
	ErrorDatabaseError      // 数据库错误
	ErrorJSONMarshalError   // JSON marshal错误
	ErrorJSONUnmarshalError // JSON Unmarshal错误
	ErrorAuthFailed         // 鉴权失败
	ErrorParamNotSupport    // 不支持的参数
)

var errorText = map[int]string{
	ErrorParamInvalid:       "Param Invalid",
	ErrorResourceNotExist:   "Resource Not Exist",
	ErrorResourceIsExist:    "ErrorResource Is Exist",
	ErrorCapNotSupport:      "cap Not Support",
	ErrorDatabaseError:      "Database Error",
	ErrorJSONMarshalError:   "JSON Marshal Error",
	ErrorJSONUnmarshalError: "JSON Unmarshal Error",
	ErrorAuthFailed:         "Authentication Failure",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func ErrorText(code int) string {
	return errorText[code]
}

func ErrorCodeList() {
	for i, j := range errorText {
		fmt.Printf("%d:%s\n", i, j)
	}
}
func ErrorCodeNew(code int, msg string) *ErrorCode {
	util.CallStack()

	return &ErrorCode{code, msg}
}

// 不具体说明message时使用默认的
func ErrorCodeDefault(code int) *ErrorCode {
	util.CallStack()

	return &ErrorCode{code, ErrorText(code)}
}

func (e ErrorCode) Error() ErrorCode {
	return e
}
