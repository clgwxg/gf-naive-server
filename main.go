package main

import (
	_ "gf-admin/internal/logic"
	_ "gf-admin/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"gf-admin/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
