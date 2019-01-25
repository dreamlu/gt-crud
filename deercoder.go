package deercoder

// db database
type DataBase interface {
	// nothing
}

// common crud
type DBCruder interface {
	// crud and search id
	Create(args map[string][]string) interface{}      // create
	Update(args map[string][]string) interface{}      // update
	Delete(id string) interface{}                     // delete
	GetBySearch(args map[string][]string) interface{} // search
	GetByID(id string) interface{}                    // by id
}
