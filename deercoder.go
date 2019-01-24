package deercoder

// db database
type DataBase interface {
	//nothing
}

// common crud
type DBCrud interface {
	// crud and search id
	Create(args map[string][]string) interface{}
	Update(args map[string][]string) interface{}
	Delete(id string) interface{}
	GetBySearch(args map[string][]string) interface{}
	GetByID(id string) interface{}
}
