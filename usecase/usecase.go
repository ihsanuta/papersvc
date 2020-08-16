package usecase

import (
	"database/sql"
	"papersvc/config"
	"papersvc/domain"
)

type Usecase struct {
	Domain domain.Domain
}

var DB *sql.DB

func init() {
	DB = config.DBInit()
}
