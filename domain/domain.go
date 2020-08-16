package domain

import (
	"database/sql"
	"papersvc/config"

	_ "github.com/go-sql-driver/mysql"
)

type Domain struct {
}

var DB *sql.DB

func init() {
	DB = config.DBInit()
}
