package cache

import (
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

var cache = gcache.New()

func cacheKey(id int) string {
	return "admin-" + gconv.String(id)
}
func SetAdminPermission(ctx context.Context, adminId int, permission []string) {
	_ = cache.Set(ctx, cacheKey(adminId), permission, 0)
}
func CancelAdminPermission(ctx context.Context, adminId int) {
	_, _ = cache.Remove(ctx, cacheKey(adminId))
}
func GetAdminPermission(ctx context.Context, adminId int) ([]string, error) {
	admin, err := cache.Get(ctx, cacheKey(adminId))
	if err != nil {
		return nil, err
	}
	var permission []string
	err = admin.Scan(&permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}
