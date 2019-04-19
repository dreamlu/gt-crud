// author:  dreamlu
package deercoder

import "github.com/dreamlu/deercoder-gin/util/lib"

// implement DBCrud
// form data
type DbCrud struct {
	// attributes
	InnerTables []string    // inner join tables
	LeftTables  []string    // left join tables
	Table       string      // table name
	Model       interface{} // table model, like User{}
	ModelData   interface{} // table model data, like var user User{}, it is 'user'

	// pager info
	ClientPage  int64       // page number
	EveryPage   int64       // Number of pages per page
}

// create
func (c *DbCrud) Create(params map[string][]string) lib.MapData {

	return CreateData(c.Table, params)
}

// create res insert id
func (c *DbCrud) CreateResID(params map[string][]string) lib.GetInfo {

	return CreateDataResID(c.Table, params)
}

// update
func (c *DbCrud) Update(params map[string][]string) lib.MapData {

	return UpdateData(c.Table, params)
}

// delete
func (c *DbCrud) Delete(id string) lib.MapData {

	return DeleteDataByName(c.Table, "id", id)
}

// search
// pager info
// clientPage : default 1
// everyPage : default 10
func (c *DbCrud) GetBySearch(params map[string][]string) lib.GetInfoPager {

	return GetDataBySearch(c.Model, c.ModelData, c.Table, params)
}

// by id
func (c *DbCrud) GetByID(id string) lib.GetInfo {

	//DB.AutoMigrate(&c.Model)
	return GetDataByID(c.ModelData, id)
}

// the same as search
// more tables
func (c *DbCrud) GetMoreBySearch(params map[string][]string) lib.GetInfoPager {

	return GetMoreDataBySearch(c.Model, c.ModelData, params, c.InnerTables, c.LeftTables)
}

// common sql
// through sql get data
func (c *DbCrud) GetDataBySQL(sql string, args ...interface{}) lib.GetInfo {

	return GetDataBySQL(c.ModelData, sql, args[:]...)
}

// common sql
// through sql get data
// args not include limit ?, ?
// args is sql and sqlnolimit common params
func (c *DbCrud) GetDataBySearchSQL(sql, sqlnolimit string, args ...interface{}) lib.GetInfoPager {

	return GetDataBySQLSearch(c.ModelData, sql, sqlnolimit, c.ClientPage, c.EveryPage, args)
}

// delete by sql
func (c *DbCrud) DeleteBySQL(sql string, args ...interface{}) lib.MapData {

	return DeleteDataBySQL(sql, args[:]...)
}

// update by sql
func (c *DbCrud) UpdateBySQL(sql string, args ...interface{}) lib.MapData {

	return UpdateDataBySQL(sql, args[:]...)
}

// create by sql
func (c *DbCrud) CreateBySQL(sql string, args ...interface{}) lib.MapData {

	return CreateDataBySQL(sql, args[:]...)
}
