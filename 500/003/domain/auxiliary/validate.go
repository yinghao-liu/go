package aux

import "ddd/domain/errs"

// 依赖注入
var Validate validator

// 校验结构体参数
type validator interface {
	Struct(s interface{}) (err errs.ErrorCoder)
}
