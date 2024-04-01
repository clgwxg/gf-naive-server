package admin

import (
	"context"
	"gf-admin/internal/dao"
	"gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
	"gf-admin/utility/cache"
	"gf-admin/utility/jwt"
	"gf-admin/utility/response"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

func init() {
	service.RegisterAdmin(&sAdmin{})
}

type sAdmin struct {
}

func (s *sAdmin) Login(ctx context.Context, in model.LoginInput) (token string, err error) {
	var emp *entity.Emp
	err = dao.Emp.Ctx(ctx).Where(do.Emp{Account: in.Account}).Scan(&emp)
	if err != nil || emp == nil {
		return "", response.NewError("账号密码不匹配", nil)
	}
	password, _ := gmd5.Encrypt(in.Password)
	if password != emp.Password {
		return "", response.NewError("账号密码不匹配", nil)
	}

	hs := jwt.NewHS()
	token, err = hs.Encode(jwt.NewCliams(int64(emp.Id)))
	if err != nil {
		return "", response.NewError("系统错误，稍后重试", nil)
	}
	return token, nil
}
func (s *sAdmin) GetAdminInfo(ctx context.Context) (*model.AdminInfo, error) {
	userId := ctx.Value("UserId")
	emp, err := service.Emp().QueryById(ctx, gconv.Int(userId))
	if err != nil || emp == nil {
		return nil, response.GetErrorByCode(3000)
	}
	var adminInfo = &model.AdminInfo{
		AdminInfo: emp,
	}
	roleMenus, err := service.Role().QueryRoleMenuByRoleId(ctx, emp.RoleId)
	if err != nil {
		adminInfo.Menu = []*entity.Menu{}
	}
	menuIds := make([]int, 0)
	for _, item := range roleMenus {
		menuIds = append(menuIds, item.MenuId)
	}
	var menus []*entity.Menu
	err = dao.Menu.Ctx(ctx).WhereIn("id", menuIds).Where(do.Menu{Status: 1}).Scan(&menus)
	if err != nil {
		adminInfo.Menu = []*entity.Menu{}
	} else {
		adminInfo.Menu = menus
	}
	if len(adminInfo.Menu) > 0 {
		permission := make([]string, 0)
		for _, item := range adminInfo.Menu {
			permission = append(permission, item.Perms)
		}
		cache.SetAdminPermission(ctx, adminInfo.AdminInfo.Id, permission)
	}

	return adminInfo, nil
}

func (s *sAdmin) Logout(ctx context.Context) error {
	userId := ctx.Value("UserId")
	cache.CancelAdminPermission(ctx, gconv.Int(userId))
	return nil

}
