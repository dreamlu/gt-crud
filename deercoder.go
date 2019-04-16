// author:  dreamlu
package deercoder

const Version = "1.2.5"

// db database
type DataBase interface {
	// nothing
}

// crud
type Crud interface {
	// crud method

	// common sql data
	// through sql, get the data
	GetDataBySQL(sql string, args ...interface{}) interface{} // single data
	// page limit ?,?
	// args not include limit ?,?
	GetDataBySearchSQL(sql, sqlnolimit string, args ...interface{}) interface{} // more data
	DeleteBySQL(sql string, args ...interface{}) interface{}
	UpdateBySQL(sql string, args ...interface{}) interface{}
	CreateBySQL(sql string, args ...interface{}) interface{}
}

// common crud
// detail impl, ==>DbCrud, implement DBCrud
// form data
type DBCruder interface {
	// crud and search id
	Create(args map[string][]string) interface{}          // create
	Update(args map[string][]string) interface{}          // update
	Delete(id string) interface{}                         // delete
	GetBySearch(args map[string][]string) interface{}     // search
	GetByID(id string) interface{}                        // by id
	GetMoreBySearch(args map[string][]string) interface{} // more search

	// common sql data
	Crud
}

// common crud
// json data
type DBCrudJer interface {
	// crud and search id
	Create(data interface{}) interface{}          // create
	Update(data interface{}) interface{}          // update
	Delete(id string) interface{}                         // delete

	// get url params
	// like form data
	GetBySearch(args map[string][]string) interface{}     // search
	GetByID(id string) interface{}                        // by id
	GetMoreBySearch(args map[string][]string) interface{} // more search

	// common sql data
	Crud
}