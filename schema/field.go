package schema

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// Field 表结构
type Field struct {
	Name    string `json:"fieldName"`
	Type    string `json:"type"`
	Comment string `json:"commnet"`
}

// GetFields 得到字段
func (f *Field) GetFields(db *gorm.DB, table string) []Field {
	var sql string = fmt.Sprintf("select COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE from information_schema.COLUMNS where table_name = '%s' and table_schema='test'", table)
	var data []Field
	var isNullAble string

	rows, err := db.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		panic("查询所有的字段失败，\n-------------------------------\n" + err.Error())
	}

	for rows.Next() {
		var temp *Field = new(Field)
		rows.Scan(&temp.Name, &temp.Type, &temp.Comment, &isNullAble)
		f.mysqlTypeToGo(&temp.Type, isNullAble)
		data = append(data, *temp)
	}

	return data
}

// MysqlTypeToGo 将数据库的类型转为go中的类型
func (f *Field) mysqlTypeToGo(types *string, isNullAble string) {
	// typeToGo 将mysql类型转为go中类型
	switch *types {
	case "int":
		*types = f.nullType("int64", isNullAble)
	case "varchar":
		*types = f.nullType("string", isNullAble)
	}
}

// NullType 得到sql中的null类型
func (f *Field) nullType(types string, isNullAble string) string {
	if strings.ToLower(isNullAble) == "no" {
		return types
	}
	return "sql.Null" + strings.Title(types)
}
