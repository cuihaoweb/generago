package codegenerate

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

// MysqlField 表结构
type MysqlField struct {
	Name    string `json:"fieldName"`
	Type    string `json:"type"`
	Comment string `json:"commnet"`
}
