// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmpDao is the data access object for table emp.
type EmpDao struct {
	table   string     // table is the underlying table name of the DAO.
	group   string     // group is the database configuration group name of current DAO.
	columns EmpColumns // columns contains all the column names of Table for convenient usage.
}

// EmpColumns defines and stores column names for table emp.
type EmpColumns struct {
	Id        string //
	Account   string // 账号
	Password  string // 密码
	NickName  string // 用户昵称
	DeptId    string // 部门id
	RoleId    string // 用户角色
	PostId    string // 用户岗位
	Phone     string // 手机号
	Email     string // 邮箱
	Remark    string // 备注
	Status    string // 用户状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// empColumns holds the columns for table emp.
var empColumns = EmpColumns{
	Id:        "id",
	Account:   "account",
	Password:  "password",
	NickName:  "nick_name",
	DeptId:    "dept_id",
	RoleId:    "role_id",
	PostId:    "post_id",
	Phone:     "phone",
	Email:     "email",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewEmpDao creates and returns a new DAO object for table data access.
func NewEmpDao() *EmpDao {
	return &EmpDao{
		group:   "default",
		table:   "emp",
		columns: empColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EmpDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EmpDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EmpDao) Columns() EmpColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EmpDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EmpDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EmpDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
