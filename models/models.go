package models

import (
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/reflect"
	"github.com/dreamlu/gt/tool/result"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/type/time"
)

type IDCom struct {
	ID uint64 `gorm:"type:bigint(20) AUTO_INCREMENT;primaryKey;" json:"id"`
}

// 通用模型
type ModelCom struct {
	IDCom
	CreateTime time.CTime `gorm:"type:datetime;DEFAULT:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间自动生成
}

// 多账号, 如果不需要多账号, 注释掉AdminID即可(ps: 为了简化部署,可直接在util/db/db.go中加入初始化appid applet账号信息)
// 账号关联
type AdminCom struct {
	ModelCom
	AdminID uint64 `json:"admin_id" gorm:"type:bigint(20);INDEX:查询索引admin_id"`
}

// ================ common ============

// CrudService service
type CrudService struct {
	GetService
	SearchService
	DeleteService
	UpdateService
	CreateService
}

type CrudServiceParam func(*CrudService)

func Get(GetService GetService) CrudServiceParam {

	return func(params *CrudService) {
		params.GetService = GetService
	}
}

func NewService(params ...CrudServiceParam) CrudService {
	param := &CrudService{}

	for _, p := range params {
		p(param)
	}
	return *param
}

func Search(SearchService SearchService) CrudServiceParam {

	return func(params *CrudService) {
		params.SearchService = SearchService
	}
}

func Delete(DeleteService DeleteService) CrudServiceParam {

	return func(params *CrudService) {
		params.DeleteService = DeleteService
	}
}

func Update(UpdateService UpdateService) CrudServiceParam {

	return func(params *CrudService) {
		params.UpdateService = UpdateService
	}
}

func Create(CreateService CreateService) CrudServiceParam {

	return func(params *CrudService) {
		params.CreateService = CreateService
	}
}

// common crud
type Com struct {
	Model interface{}
	CrudService
}

// get data, by id
func (c *Com) Get(params cmap.CMap) (data interface{}, err error) {

	if c.GetService != nil {
		return c.GetService.Get(params)
	}
	data = reflect.New(c.Model)
	crud := gt.NewCrud(gt.Model(c.Model), gt.Data(data))
	if err = crud.Get(params).Error(); err != nil {
		return
	}
	return
}

// get data, limit and search
func (c *Com) Search(params cmap.CMap) (datas interface{}, pager result.Pager, err error) {

	if c.SearchService != nil {
		return c.SearchService.Search(params)
	}
	datas = reflect.NewArray(c.Model)
	crud := gt.NewCrud(gt.Model(c.Model), gt.Data(datas))
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		return nil, pager, cd.Error()
	}
	return datas, cd.Pager(), nil
}

// delete data, by id
func (c *Com) Delete(id interface{}) error {

	if c.DeleteService != nil {
		return c.DeleteService.Delete(id)
	}
	return gt.NewCrud(gt.Model(c.Model)).Delete(id).Error()
}

// update data
func (c *Com) Update(data interface{}) error {

	if c.UpdateService != nil {
		return c.UpdateService.Update(data)
	}
	crud := gt.NewCrud(gt.Model(c.Model), gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	return nil
}

// create data
func (c *Com) Create(data interface{}) error {

	if c.CreateService != nil {
		return c.CreateService.Create(data)
	}
	crud := gt.NewCrud(gt.Model(c.Model), gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		return err
	}
	return nil
}

func (c *Com) M() interface{} {

	return c.Model
}
