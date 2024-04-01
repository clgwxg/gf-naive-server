package middleware

import (
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Code int         `json:"code"    dc:"Error code"`
	Msg  string      `json:"msg" dc:"Error message"`
	Data interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

func MiddlewareResponse(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = response.CodeOK
		}
	}
	r.Response.WriteJson(DefaultHandlerResponse{
		Code: code.Code(),
		Msg:  msg,
		Data: res,
	})
}
