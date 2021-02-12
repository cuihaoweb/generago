package main

import (
	"github.com/cuihaoweb/generago/schema"
	"github.com/cuihaoweb/generago/templates"
	"github.com/cuihaoweb/generago/util"
)

// DataSource xx
type DataSource schema.DataSource

var dataSource = schema.DataSource{}
var table = schema.Table{}
var field = schema.Field{}
var template = schema.Template{}
var templateStruct = schema.TemplateStruct{}
var fileMode = schema.ModelFile{}

// SetDataSource xx
func SetDataSource(data DataSource) {
	util.CopyStruct(data, &dataSource)
	dataSource.DefaultDataSource()
}

// SetOutDir xx
func SetOutDir(dirName string) {
	fileMode.DirName = dirName
}

// Execute xx
func Execute() {
	var db = dataSource.GetMysqlConn()

	template.TemplateHandler(templates.MODEL)

	var tables = table.GetTables(db)
	for _, val := range tables {
		fields := field.GetFields(db, val)
		str := templateStruct.GetContent(template.Temp, val, fields)
		fileMode.OutputFile(val, str)
	}
}
func main() {
	SetDataSource(DataSource{
		DbName:   "test",
		User:     "root",
		Password: "ch1997",
	})
	SetOutDir("./model")
	Execute()
}
