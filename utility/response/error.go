package response

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var mapErr map[int]string = map[int]string{
	3000: "权限不足",
}

func NewError(msg string, code *int) error {
	_code := failCode
	if code != nil {
		_code = *code
	}
	return gerror.NewWithOption(gerror.Option{
		Stack: false,
		Text:  msg,
		Code:  gcode.New(_code, "", nil),
	})
}
func GetErrorByCode(code int) error {
	return NewError(mapErr[code], &code)
}
