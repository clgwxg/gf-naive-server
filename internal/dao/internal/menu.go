// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MenuDao is the data access object for table menu.
type MenuDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns MenuColumns // columns contains all the column names of Table for convenient usage.
}

// MenuColumns defines and stores column names for table menu.
type MenuColumns struct {
	Id        string //
	Name      string // 菜单名称
	ParentId  string // 父级id
	Type      string // 菜单类型：1目录2菜单3按钮
	Icon      string // 菜单图标
	Url       string // 路由地址
	IsFrame   string // 是否外链：1是0否
	Path      string // 组件路径
	Query     string // 路由参数
	IsCache   string // 是否缓存：1缓存0不缓存
	Perms     string // 权限字符
	Visible   string // 是否显示：1显示0不显示
	Status    string // 菜单状态：1正常0禁用
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// menuColumns holds the columns for table menu.
var menuColumns = MenuColumns{
	Id:        "id",
	Name:      "name",
	ParentId:  "parent_id",
	Type:      "type",
	Icon:      "icon",
	Url:       "url",
	IsFrame:   "is_frame",
	Path:      "path",
	Query:     "query",
	IsCache:   "is_cache",
	Perms:     "perms",
	Visible:   "visible",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewMenuDao creates and returns a new DAO object for table data access.
func NewMenuDao() *MenuDao {
	return &MenuDao{
		group:   "default",
		table:   "menu",
		columns: menuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MenuDao) Columns() MenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
