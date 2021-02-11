package schema

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlConf 用户配置
type MysqlConf struct {
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	DbName    string `json:"dbName"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parseTime"`
	Loc       string `json:"loc"`
}

// GetTables 获取数据中所有表
func (m *MysqlConf) GetTables() []string {
	var sql string = "show tables"
	var data []string

	db := GetMysqlConn(*m)
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

// GetFields 得到字段
func (m *MysqlConf) GetFields(table string) []MysqlField {
	var sql string = fmt.Sprintf("select COLUMN_NAME, DATA_TYPE, COLUMN_COMMENT, IS_NULLABLE from information_schema.COLUMNS where table_name = '%s' and table_schema='test'", table)
	var data []MysqlField
	var isNullAble string

	db := GetMysqlConn(*m)
	rows, err := db.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		panic("查询所有的字段失败，\n-------------------------------\n" + err.Error())
	}

	for rows.Next() {
		var temp *MysqlField = new(MysqlField)
		rows.Scan(&temp.Name, &temp.Type, &temp.Comment, &isNullAble)
		MysqlTypeToGo(&temp.Type, isNullAble)
		data = append(data, *temp)
	}

	return data
}

// DefaultMysqlConf mysql默认配置
func DefaultMysqlConf(conf *MysqlConf) {
	conf.IP = "127.0.0.1"
	conf.Port = 3306
	conf.Charset = "utf8mb4"
	conf.ParseTime = true
	conf.Loc = "Local"
}

// MysqlConfToString 拼接mysql的连接字符串
func MysqlConfToString(conf MysqlConf) string {
	str := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s"

	var parseTime string = strings.Title("true")
	if conf.ParseTime == false {
		parseTime = strings.Title("false")
	}

	connerStr := fmt.Sprintf(str, conf.User, conf.Password, conf.IP, conf.Port, conf.DbName, conf.Charset, parseTime, conf.Loc)

	return connerStr
}

// GetMysqlConn 得到数据库连接
func GetMysqlConn(mysqlConf MysqlConf) *gorm.DB {
	DefaultMysqlConf(&mysqlConf)
	var str string = MysqlConfToString(mysqlConf)

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，\n-------------------------------\n" + err.Error())
	}

	return db
}

// MysqlTypeToGo 将数据库的类型转为go中的类型
func MysqlTypeToGo(types *string, isNullAble string) {
	// typeToGo 将mysql类型转为go中类型
	switch *types {
	case "int":
		*types = NullType("int64", isNullAble)
	case "varchar":
		*types = NullType("string", isNullAble)
	}
}

// NullType 得到sql中的null类型
func NullType(types string, isNullAble string) string {
	if strings.ToLower(isNullAble) == "no" {
		return types
	}
	return "sql.Null" + strings.Title(types)
}

// MysqlField 表结构
type MysqlField struct {
	Name    string `json:"fieldName"`
	Type    string `json:"type"`
	Comment string `json:"commnet"`
}
