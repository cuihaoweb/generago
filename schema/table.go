package schema

import "gorm.io/gorm"

// Table xx
type Table struct{}

// GetTables 获取数据中所有表
func (t *Table) GetTables(db *gorm.DB) []string {
	var sql string = "show tables"
	var data []string

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
