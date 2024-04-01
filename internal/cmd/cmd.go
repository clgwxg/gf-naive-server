package cmd

import (
	"context"
	"gf-admin/internal/controller/admin"
	"gf-admin/internal/controller/dept"
	"gf-admin/internal/controller/emp"
	"gf-admin/internal/controller/menu"
	"gf-admin/internal/controller/post"
	"gf-admin/internal/controller/role"
	"gf-admin/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.MiddlewareResponse, middleware.Auth)
				group.Group("admin", func(group *ghttp.RouterGroup) {
					group.Bind(
						admin.NewV1(),
						menu.NewV1(),
						role.NewV1(),
						post.NewV1(),
						dept.NewV1(),
						emp.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
