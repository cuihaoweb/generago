package generatego

import (
	"fmt"
	"strings"
)

// 拼接mysql的连接字符串
func mysqlConfToString(conf MysqlConf) string {
	str := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s"

	var parseTime string = strings.Title("true")
	if conf.ParseTime == false {
		parseTime = strings.Title("false")
	}

	connerStr := fmt.Sprintf(str, conf.User, conf.Password, conf.IP, conf.Port, conf.DbName, conf.Charset, parseTime, conf.Loc)

	return connerStr
}
