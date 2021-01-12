package reflect

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"testing"
)

type Client struct {
	Name string `gorm:"type:varchar(30);" json:"name" valid:"required,len=2-20"` // 昵称
}

func TestNew(t *testing.T) {
	//var ct Client
	t.Log(Client{})
	t.Log(New(Client{Name: "test"}))
	t.Log(New([]*Client{}))
	t.Log(NewArray(Client{}))
}

func TestAnnotion(t *testing.T) {
	fset := token.NewFileSet()
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	path, _ := filepath.Abs("./reflect.go")
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	// 打印语法树
	//ast.Print(fset, f)
	//fset.Iterate(f)
	for _, v := range cmap.Comments() {
		t.Log(v.Text())
	}

	fset = token.NewFileSet()
	f, err = parser.ParseFile(fset, "reflect.go", nil, 0)

	if err != nil {
		panic(err)
	}

	// hard coding looking these up
	//typeDecl := f.Decls[1].(*ast.FuncDecl)
	//structDecl := typeDecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	//fields := structDecl.Fields.List

	//for _, field := range fields {
	//	//typeExpr := field.Type
	//
	//	//start := typeExpr.Pos() - 1
	//	//end := typeExpr.End() - 1
	//
	//	// grab it in source
	//	//typeInSource := src[start:end]
	//
	//	//fmt.Println(typeInSource)
	//	t.Log(field.Type)
	//	t.Log(field.Names[0].Obj)
	//}

}
