package generatego

import (
	"fmt"
	"text/template"
)

var mysqlConfig MysqlConf                       // mysqlConfig 数据库配置信息
var temp *template.Template                     // 模板
var databaseMap = make(map[string][]MysqlField) // 保存的数据库数据
var outDir string                               // 输出的目录

// SetDataSource 设置数据库的连接信息
func SetDataSource(mysqlConf MysqlConf) {
	verifyDataSource(mysqlConf)

	mysqlConfig = mysqlConf

	arrList1 := getTables()
	for _, value := range arrList1 {
		databaseMap[value] = getFields(value)
	}
}
func verifyDataSource(mysqlConf MysqlConf) {
	// verifyDataSource 验证数据库配置
	if mysqlConf.DbName == "" {
		panic("数据库配置\t=>\t数据库名DbName不能为空\n------------------------------------------")
	} else if mysqlConf.User == "" {
		panic("数据库配置\t=>\t数据库用户User不能为空\n------------------------------------------")
	} else if mysqlConf.Password == "" {
		panic("数据库配置\t=>\t数据库密码Password不能为空\n------------------------------------------")
	}
}

// SetOutDir 设置输出路径
func SetOutDir(dir string) {
	var len = len(dir)

	if dir[len-1] == '/' {
		outDir = dir
		return
	}
	outDir = dir + "/"
}

// Execute 执行操作
func Execute() {
	initTemplate()

	templateHandler()
}

func templateHandler() {
	// 模板处理部分
	for key, val := range databaseMap {
		str := getTemplateContent(key, val)
		fmt.Print(str)
		createFile(outDir+key, str)
	}
}

// // InitTest 测试
// func InitTest() {
// 	SetDataSource(MysqlConf{DbName: "test", User: "root", Password: "ch1997"})
// 	SetOutDir("./model")
// 	initTemplate()
// 	templateHandler()
// }
