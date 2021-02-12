package schema

import (
	"strings"
	"text/template"
)

// Template xx
type Template struct {
	template.Template
	Temp *template.Template
}

func (t *Template) newTemplate() *Template {
	t.Temp = t.New("generago")
	return t
}

func (t *Template) addFuncs() *Template {
	temp := t.Temp

	temp.Funcs(template.FuncMap{
		"FuncTitle": func(str string) string {
			if str == "id" {
				return "ID"
			}
			return strings.Title(str)
		},
		"FuncComment": func(str string) string {
			if str != "" {
				return "//" + str
			}
			return ""
		},
	})

	return t
}

func (t *Template) parseStr(str string) *Template {
	var temp = t.Temp

	temp, err := temp.Parse(str)
	if err != nil {
		panic("解析模板文件失败\n--------------------------------------------")
	}

	return t
}

// TemplateHandler xx
func (t *Template) TemplateHandler(str string) {
	t.newTemplate().addFuncs().parseStr(str)
}
