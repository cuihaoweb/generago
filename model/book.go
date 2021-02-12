package model

import "database/sql"

// Book xx
type Book struct {
    Name 	 sql.NullString 	 `json:"name"` 	 
    Pirce 	 sql.NullString 	 `json:"pirce"` 	 
    
}