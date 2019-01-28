// author:  dreamlu
package deercoder

const Version = "1.2.0"

// db database
type DataBase interface {
	// nothing
}

// common crud
// detail impl, ==>DbCrud, implement DBCrud
type DBCruder interface {
	// crud and search id
	Create(args map[string][]string) interface{}          // create
	Update(args map[string][]string) interface{}          // update
	Delete(id string) interface{}                         // delete
	GetBySearch(args map[string][]string) interface{}     // search
	GetByID(id string) interface{}                        // by id
	GetMoreBySearch(args map[string][]string) interface{} // more search

	// through sql, get the data
	GetDataBySQL(sql string, args ...string) interface{}                   // single data
	// page limit ?, ?
	GetDataBySearchSQL(sql, sqlnolimit string, args ...string) interface{} // more data
}
