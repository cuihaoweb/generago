package schema

import (
	"bytes"
	"text/template"
)

// TemplateStruct xx
type TemplateStruct struct {
	TableName string  `json:"tableName"`
	FieldList []Field `json:"fieldList"`
}

// GetContent xx
func (t *TemplateStruct) GetContent(temp *template.Template, key string, val []Field) string {
	var buf bytes.Buffer
	var mapContent = TemplateStruct{key, val}

	if err := temp.Execute(&buf, mapContent); err != nil {
		panic("读取模板内容失败\n--------------------------------------------")
	}

	return string(buf.Bytes())
}
