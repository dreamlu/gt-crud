package models

import (
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
)

type Service interface {
	GetService
	SearchService
	DeleteService
	UpdateService
	CreateService
	M() interface{}
}

type GetService interface {
	Get(params cmap.CMap) (data interface{}, err error)
}

type SearchService interface {
	Search(params cmap.CMap) (datas interface{}, pager result.Pager, err error)
}

type DeleteService interface {
	Delete(id interface{}) error
}

type UpdateService interface {
	Update(data interface{}) error
}

type CreateService interface {
	Create(data interface{}) error
}
