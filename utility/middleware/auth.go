package middleware

import (
	"gf-admin/utility/cache"
	"gf-admin/utility/jwt"
	"gf-admin/utility/tools"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmeta"
	"reflect"
)

func Auth(r *ghttp.Request) {
	//权限标识
	permission := ""
	auth := false
	handler := r.GetServeHandler().Handler
	var objectReq reflect.Value
	if handler.Info.Type != nil && handler.Info.Type.NumIn() == 2 {
		objectReq = reflect.New(handler.Info.Type.In(1))
		if v := gmeta.Get(objectReq, "permission"); !v.IsEmpty() {
			permission = v.String()
		}
		if v := gmeta.Get(objectReq, "auth"); !v.IsEmpty() {
			auth = v.Bool()
		}
	}
	if auth == true || permission != "" {
		authorization := r.GetHeader("Authorization")
		hs := jwt.NewHS()
		var cliams jwt.SignCliams
		err := hs.Decode(authorization, &cliams)
		if err == nil && cliams.UserId != 0 {
			r.SetCtxVar("UserId", cliams.UserId)
			if auth == true {
				r.Middleware.Next()
				return
			} else {
				permissions, err := cache.GetAdminPermission(r.Context(), int(cliams.UserId))
				if err == nil && tools.Contain(permissions, permission) {
					r.Middleware.Next()
					return
				} else {
					if v := gmeta.Get(objectReq, "summary"); !v.IsEmpty() {
						summary := v.String()
						r.Response.WriteJson(DefaultHandlerResponse{
							Code: 3001,
							Msg:  "没有" + summary + "权限,请联系管理员",
						})
						return
					}
				}
			}
		}
		r.Response.WriteJson(DefaultHandlerResponse{
			Code: 3000,
			Msg:  "权限不足",
			Data: nil,
		})
	} else {
		r.Middleware.Next()
	}
}
