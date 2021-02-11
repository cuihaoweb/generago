package templates

// MODEL model字符模板
const MODEL = `package model

import "database/sql"

// {{FuncTitle .TableName}} xx
type {{FuncTitle .TableName}} struct {
    {{range $k, $v := .FieldList}}` +
	"{{FuncTitle $v.Name}} \t {{$v.Type}} \t `json:\"{{$v.Name}}\"` \t {{FuncComment $v.Comment}}" + `
    {{end}}
}`
