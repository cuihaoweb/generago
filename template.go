package codegenerate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var str1 string = `package model

import "database/sql"

// {{FuncTitle .TableName}} xx
type {{FuncTitle .TableName}} struct {
    {{range $k, $v := .FieldList}}` +
	"{{FuncTitle $v.Name}} \t {{$v.Type}} \t `json:\"{{$v.Name}}\"` \t {{FuncComment $v.Comment}}" + `
    {{end}}
}`

// initTemplate 初始化
func initTemplate() {
	var err error

	temp = template.New("code-generate")

	temp.Funcs(template.FuncMap{
		"FuncTitle": func(str string) string {
			return strings.Title(str)
		},
		"FuncComment": func(str string) string {
			if str != "" {
				return "//" + str
			}
			return ""
		},
	})
	// temp, err = temp.ParseFiles("./temp/model.tmpl")
	temp, err = temp.Parse(str1)
	if err != nil {
		panic("解析模板文件失败\n--------------------------------------------")
	}
}

func getTemplateContent(key string, val []MysqlField) string {
	var buf bytes.Buffer
	var mapContent = make(map[string]interface{})

	mapContent["TableName"] = key
	mapContent["FieldList"] = val

	if err := temp.Execute(&buf, mapContent); err != nil {
		panic("读取模板内容失败\n--------------------------------------------")
	}

	return string(buf.Bytes())
}

func createFile(filedName string, content string) {
	var path = filedName + ".go"
	fmt.Print(path)

	//创建目录
	os.MkdirAll("./model", os.ModePerm)

	if err := ioutil.WriteFile(path, []byte(content), 0666); err != nil {
		panic("写入文件失败\n--------------------------------------------")
	}
}
