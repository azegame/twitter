package repository

import (
	"database/sql"
	
	"twitter/model/connection"
)

var	db *sql.DB

func init() {
	db = dao.GetDB()
}

