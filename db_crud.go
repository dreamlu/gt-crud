package deercoder

// implment DBCrud
type DbCrudImpl struct {
	// nothing
}

func (db *DbCrudImpl) Create(args map[string][]string) interface{} {

	return nil
}

func (db *DbCrudImpl) Update(args map[string][]string) interface{} {

	return nil
}

func (db *DbCrudImpl) Delete(id string) interface{} {

	return nil
}

func (db *DbCrudImpl) GetBySearch(args map[string][]string) interface{} {

	return nil//CreateData("user", args)
}

func (db *DbCrudImpl) GetByID(id string) interface{} {

	//DB.AutoMigrate(&User{})
	//var user = User{}
	//return GetDataById(&user, id)
	return nil
}
