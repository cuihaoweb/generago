package model

import "database/sql"

// User xx
type User struct {
    ID 	 int64 	 `json:"id"` 	 //用户id
    Uname 	 sql.NullString 	 `json:"uname"` 	 
    Age 	 sql.NullInt64 	 `json:"age"` 	 
    
}