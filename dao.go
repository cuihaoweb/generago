package generatego

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// defaultMysqlConf mysql默认配置
func defaultMysqlConf(conf *MysqlConf) {
	conf.IP = "127.0.0.1"
	conf.Port = 3306
	conf.Charset = "utf8mb4"
	conf.ParseTime = true
	conf.Loc = "Local"
}

// getMysqlConn 得到数据库连接
func getMysqlConn(mysqlConf MysqlConf) *gorm.DB {
	defaultMysqlConf(&mysqlConf)
	var str string = mysqlConfToString(mysqlConf)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，\n-------------------------------\n" + err.Error())
	}

	return db
}

// getTables 获取数据中所有表
func getTables() []string {
	var sql string = "show tables"
	var data []string

	db := getMysqlConn(mysqlConfig)
	rows, err := db.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		panic("查询所有的表失败，\n-------------------------------\n" + err.Error())
	}

	var temp string
	for rows.Next() {
		rows.Scan(&temp)
		data = append(data, temp)
	}

	return data
}

// getFields 得到字段
func getFields(table string) []MysqlField {
	var sql string = fmt.Sprintf("select COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE from information_schema.COLUMNS where table_name = '%s' and table_schema='test'", table)
	var data []MysqlField
	var isNullAble string

	db := getMysqlConn(mysqlConfig)
	rows, err := db.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		panic("查询所有的字段失败，\n-------------------------------\n" + err.Error())
	}

	for rows.Next() {
		var temp *MysqlField = new(MysqlField)
		rows.Scan(&temp.Name, &temp.Type, &temp.Comment, &isNullAble)
		typeToGo(&temp.Type, isNullAble)
		data = append(data, *temp)
	}

	return data
}
func typeToGo(types *string, isNullAble string) {
	// typeToGo 将mysql类型转为go中类型
	switch *types {
	case "int":
		*types = nullType("int64", isNullAble)
	case "varchar":
		*types = nullType("string", isNullAble)
	}
}
func nullType(types string, isNullAble string) string {
	if strings.ToLower(isNullAble) == "no" {
		return types
	}
	return "sql.Null" + strings.Title(types)
}
