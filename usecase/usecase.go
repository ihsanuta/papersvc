package usecase

import (
	"database/sql"

	"github.com/ihsanuta/papersvc/config"
	"github.com/ihsanuta/papersvc/domain"
)

type Usecase struct {
	Domain domain.Domain
}

var DB *sql.DB

func init() {
	DB = config.DBInit()
}
