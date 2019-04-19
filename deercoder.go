// author:  dreamlu
package deercoder

import "github.com/dreamlu/deercoder-gin/util/lib"

const Version = "1.2.8"

// db database
type DataBase interface {
	// nothing
}

// crud
type Crud interface {
	// crud method

	// get url params
	// like form data
	GetBySearch(args map[string][]string) lib.GetInfoPager     // search
	GetByID(id string) lib.GetInfo                             // by id
	GetMoreBySearch(args map[string][]string) lib.GetInfoPager // more search

	// common sql data
	// through sql, get the data
	GetDataBySQL(sql string, args ...interface{}) lib.GetInfo // single data
	// page limit ?,?
	// args not include limit ?,?
	GetDataBySearchSQL(sql, sqlnolimit string, args ...interface{}) lib.GetInfoPager // more data
	DeleteBySQL(sql string, args ...interface{}) lib.MapData
	UpdateBySQL(sql string, args ...interface{}) lib.MapData
	CreateBySQL(sql string, args ...interface{}) lib.MapData
}

// common crud
// detail impl, ==>DbCrud, implement DBCrud
// form data
type DBCruder interface {
	// crud and search id
	Create(args map[string][]string) lib.MapData      // create
	CreateResID(args map[string][]string) lib.GetInfo // create res insert id
	Update(args map[string][]string) lib.MapData      // update
	Delete(id string) lib.MapData                     // delete

	// common sql data
	Crud
}

// common crud
// json data
type DBCrudJer interface {
	// crud and search id
	Create(data interface{}) lib.MapData      // create
	CreateResID(data interface{}) lib.GetInfo // create res insert id
	Update(data interface{}) lib.MapData      // update
	Delete(id string) lib.MapData             // delete

	// common sql data
	Crud
}
