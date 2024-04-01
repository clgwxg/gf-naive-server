package model

import "gf-admin/internal/model/entity"

type BaseDeptInput struct {
	Name     string
	ParentId int64
	Status   uint8
	Leader   string
	Phone    string
	Email    string
}

type DeptCreateInput struct {
	BaseDeptInput
}

type DeptQueryInput struct {
	Name   string
	Status *int8
}

type DeptTree struct {
	entity.Dept
	Children []*DeptTree `json:"children"`
}

type DeptUpdateInput struct {
	Id int
	BaseDeptInput
}
