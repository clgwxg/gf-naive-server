package model

import "gf-admin/internal/model/entity"

type LoginInput struct {
	Account  string
	Password string
}

type AdminInfo struct {
	AdminInfo *entity.Emp    `json:"adminInfo"`
	Menu      []*entity.Menu `json:"menu"`
}
