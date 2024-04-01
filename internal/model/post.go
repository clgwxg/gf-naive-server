package model

import "gf-admin/api"

type BasePostInput struct {
	Name   string
	Code   string
	Status uint8
	Remark string
}
type PostCreateInput struct {
	BasePostInput
}

type PostUpdateInput struct {
	Id int
	BasePostInput
}

type PostQueryInput struct {
	Name   string
	Code   string
	Status *uint8
	api.Page
}
