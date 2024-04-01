// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf-admin/internal/model"
	"gf-admin/internal/model/entity"
)

type (
	IPost interface {
		Create(ctx context.Context, in model.PostCreateInput) error
		Update(ctx context.Context, in model.PostUpdateInput) error
		DeleteById(ctx context.Context, id int) error
		Query(ctx context.Context, in model.PostQueryInput) ([]entity.Post, int, error)
		GetPostById(ctx context.Context, id int) (*entity.Post, error)
	}
)

var (
	localPost IPost
)

func Post() IPost {
	if localPost == nil {
		panic("implement not found for interface IPost, forgot register?")
	}
	return localPost
}

func RegisterPost(i IPost) {
	localPost = i
}
