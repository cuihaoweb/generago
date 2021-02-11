package schema

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/cuihaoweb/generago/templates"
)

// Template xx
type Template struct {
	Temp *template.Template
}

// initTemplate 初始化
func (t *Template) initTemplate() {
	var temp = (*t).Temp
	var str = templates.MODEL
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
	temp, err = temp.Parse(str)
	if err != nil {
		panic("解析模板文件失败\n--------------------------------------------")
	}
}

func (t *Template) getTemplateContent(key string, val []MysqlField) string {
	var temp = (*t).Temp
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

	//创建目录
	os.MkdirAll(path, os.ModePerm)

	if err := ioutil.WriteFile(path, []byte(content), 0666); err != nil {
		panic("写入文件失败\n--------------------------------------------")
	}
}
