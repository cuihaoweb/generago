package schema

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DataSource 用户配置
type DataSource struct {
	IP        string `json:"ip"`
	Port      int    `json:"port"`
	DbName    string `json:"dbName"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parseTime"`
	Loc       string `json:"loc"`
}

// mysqlConfToString 拼接mysql的连接字符串
func (d *DataSource) mysqlConfToString() string {
	str := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s"

	var parseTime string = strings.Title("true")
	if d.ParseTime == false {
		parseTime = strings.Title("false")
	}

	return fmt.Sprintf(str, d.User, d.Password, d.IP, d.Port, d.DbName, d.Charset, parseTime, d.Loc)
}

// GetMysqlConn 得到数据库连接
func (d *DataSource) GetMysqlConn() *gorm.DB {
	var str string = d.mysqlConfToString()

	db, err := gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败，\n-------------------------------\n" + err.Error())
	}

	return db
}

// DefaultDataSource mysql默认配置
func (d *DataSource) DefaultDataSource() {
	d.IP = "127.0.0.1"
	d.Port = 3306
	d.Charset = "utf8mb4"
	d.ParseTime = true
	d.Loc = "Local"
}
