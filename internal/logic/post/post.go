package post

import (
	"context"
	"gf-admin/internal/dao"
	"gf-admin/internal/model"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
	"gf-admin/internal/service"
	"gf-admin/utility/response"
)

type sPost struct {
}

func init() {
	service.RegisterPost(&sPost{})
}
func (s *sPost) Create(ctx context.Context, in model.PostCreateInput) error {
	var post *entity.Post
	_ = dao.Post.Ctx(ctx).Where(do.Post{Name: in.Name}).WhereOr(do.Post{
		Code: in.Code,
	}).Scan(&post)
	if post != nil {
		return response.NewError("岗位名称或岗位编码已存在", nil)
	}
	_, err := dao.Post.Ctx(ctx).Data(do.Post{
		Name:   in.Name,
		Code:   in.Code,
		Remark: in.Remark,
		Status: in.Status,
	}).Insert()
	return err
}

func (s *sPost) Update(ctx context.Context, in model.PostUpdateInput) error {
	var post *entity.Post
	_ = dao.Post.Ctx(ctx).Where(do.Post{Name: in.Name}).WhereOr(do.Post{
		Code: in.Code,
	}).Scan(&post)
	if post != nil && in.Id != post.Id {
		return response.NewError("岗位名称或岗位编码已存在", nil)
	}
	_, err := dao.Post.Ctx(ctx).WherePri(in.Id).Data(do.Post{
		Name:   in.Name,
		Code:   in.Code,
		Remark: in.Remark,
		Status: in.Status,
	}).Update()
	return err
}

func (s *sPost) DeleteById(ctx context.Context, id int) error {
	_, err := dao.Post.Ctx(ctx).WherePri(id).Delete()
	return err
}

func (s *sPost) Query(ctx context.Context, in model.PostQueryInput) ([]entity.Post, int, error) {
	m := dao.Post.Ctx(ctx)
	if in.Name != "" {
		m = m.WhereLike("name", "%"+in.Name+"%")
	}
	if in.Code != "" {
		m = m.WhereLike("code", "%"+in.Code+"%")
	}
	if in.Status != nil {
		m = m.Where(do.Post{Status: &in.Status})
	}
	total := 0
	var list []entity.Post
	err := m.Page(in.PageNum, in.PageSize).
		ScanAndCount(&list, &total, true)
	if err != nil {
		return nil, 0, response.NewError("查询失败，稍后重试", nil)
	}
	return list, total, nil
}

func (s *sPost) GetPostById(ctx context.Context, id int) (*entity.Post, error) {
	var post *entity.Post
	err := dao.Post.Ctx(ctx).WherePri(id).Scan(&post)
	return post, err
}
