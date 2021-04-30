package models

import (
	"demo/util/result"
	"github.com/dreamlu/gt"
	"github.com/dreamlu/gt/tool/reflect"
	"github.com/dreamlu/gt/tool/type/cmap"
	"github.com/dreamlu/gt/tool/type/time"
	"github.com/dreamlu/gt/tool/util/hump"
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

// CrudParams service
type CrudParams struct {
	trans    gt.Crud
	gtParams []gt.Param
}

type CrudServiceParam func(*CrudParams)

func NewCrudService(params ...CrudServiceParam) CrudParams {
	param := &CrudParams{}

	for _, p := range params {
		p(param)
	}
	return *param
}

// Trans transaction
func Trans(trans gt.Crud) CrudServiceParam {
	return func(params *CrudParams) {
		params.trans = trans
	}
}

// 增加gt参数
func GtParams(gtParams ...gt.Param) CrudServiceParam {
	return func(params *CrudParams) {
		params.gtParams = gtParams
	}
}

// common crud
type Com struct {
	Model interface{}
	CrudParams
}

func NewService(model interface{}, params ...CrudServiceParam) *Com {
	return &Com{
		Model:      model,
		CrudParams: NewCrudService(params...),
	}
}

// get data, by id
func (c *Com) GetByID(id interface{}) (data interface{}, err error) {

	data = reflect.New(c.Model)
	crud := c.Crud().Params(gt.Model(c.Model), gt.Data(data)).Params(c.gtParams...)
	if err = crud.GetByID(id).Error(); err != nil {
		return
	}
	return
}

// get data, by id
func (c *Com) Get(params cmap.CMap) (data interface{}, err error) {

	data = reflect.New(c.Model)
	crud := c.Crud().Params(gt.Model(c.Model), gt.Data(data)).Params(c.gtParams...)
	if err = crud.Get(params).Error(); err != nil {
		return
	}
	return
}

// get data, limit and search
func (c *Com) Search(params cmap.CMap) (datas interface{}, pager result.Pager, err error) {

	datas = reflect.NewArray(c.Model)
	crud := c.Crud().Params(gt.Model(c.Model), gt.Data(datas)).Params(c.gtParams...)
	cd := crud.GetBySearch(params)
	if cd.Error() != nil {
		return nil, pager, cd.Error()
	}
	pager.Pager = cd.Pager()
	return datas, pager, nil
}

// delete data, by id
func (c *Com) Delete(id interface{}) error {

	return c.Crud().Params(gt.Model(c.Model)).Delete(id).Error()
}

// update data
func (c *Com) Update(data interface{}) error {

	crud := c.Crud().Params(gt.Model(c.Model), gt.Data(data))
	if err := crud.Update().Error(); err != nil {
		//log.Log.Error(err.Error())
		return err
	}
	return nil
}

// create data
func (c *Com) Create(data interface{}) error {

	crud := c.Crud().Params(gt.Model(c.Model), gt.Data(data))
	if err := crud.Create().Error(); err != nil {
		return err
	}
	return nil
}

func (c *Com) M() interface{} {

	return c.Model
}

func (c *Com) Crud() gt.Crud {
	var cd = gt.NewCrud()
	if c.trans != nil {
		cd = c.trans
		cd.Params(gt.Table(hump.HumpToLine(reflect.StructName(c.Model))))
	}
	return cd
}
