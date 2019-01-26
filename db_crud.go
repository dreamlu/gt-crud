// author:  dreamlu
package deercoder

// implment DBCrud
type DbCrud struct {
	// attributes
	Tables    string      // table name
	Model     interface{} // table model, like User{}
	ModelData interface{} // table model data, like var user User{}, it is 'user'
}

// create
func (c *DbCrud) Create(args map[string][]string) interface{} {

	return CreateData(c.Tables, args)
}
// update
func (c *DbCrud) Update(args map[string][]string) interface{} {

	return UpdateData(c.Tables, args)
}

// delete
func (c *DbCrud) Delete(id string) interface{} {

	return DeleteDataByName(c.Tables, "id", id)
}

// search
// pager info
// clientPage : default 1
// everyPage : default 10
func (c *DbCrud) GetBySearch(args map[string][]string) interface{} {

	return GetDataBySearch(c.Model, c.ModelData, c.Tables, args)
}

// by id
func (c *DbCrud) GetByID(id string) interface{} {

	//DB.AutoMigrate(&c.Model)
	return GetDataById(c.ModelData, id)
}
