package domain

import (
	"database/sql"
	"papersvc/entity"
	x "papersvc/errors"

	"github.com/gin-gonic/gin"
)

func (d *Domain) CreateSQLAccount(c *gin.Context, v entity.Account) (entity.Account, error) {
	row, err := DB.Exec("INSERT INTO accounts (username, password, created_at) VALUES (?,?,?)",
		v.Username,
		v.Password,
		v.CreatedAt)
	if err != nil {

		return v, err
	}

	v.ID, err = row.LastInsertId()
	if err != nil {
		return v, err
	}

	return v, nil
}

func (d *Domain) LoginSQLAccount(c *gin.Context, v entity.ReqAccount) (entity.Account, error) {
	var result entity.Account
	row := DB.QueryRow("SELECT id FROM accounts WHERE username = ? AND password = ?",
		v.Username,
		v.Password)

	err := row.Scan(
		&result.ID,
	)
	if err == sql.ErrNoRows {
		return result, x.WrapWithCode(err, x.CodeHTTPNotFound, "GetSQLFinancialTrxByID Not Found")
		// return result,
	} else if err != nil {
		return result, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancialTrxByID")
	}

	return result, nil
}
