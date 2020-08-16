package domain

import (
	"database/sql"
	"log"
	"papersvc/entity"
	x "papersvc/errors"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Domain) CreateSQLFinancialTrx(c *gin.Context, v entity.FinancialTrx) (entity.FinancialTrx, error) {
	row, err := DB.Exec("INSERT INTO finance_trx(acc_id, created_at, updated_at) VALUES (?,?,?);",
		v.Account.ID,
		v.CreatedAt,
		v.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return v, x.WrapWithCode(err, x.CodeSQLCreate, "CreateSQLFinancialTrx")
	}

	v.ID, err = row.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return v, x.WrapWithCode(err, x.CodeSQLCannotRetrieveLastInsertID, `ErrorGetFinancialTrxID`)
	}

	return v, nil
}

func (d *Domain) GetSQLFinancialTrxByID(c *gin.Context, vid int64) (entity.FinancialTrx, error) {
	var result entity.FinancialTrx
	row := DB.QueryRow("SELECT finance_trx.id, finance_trx.created_at, finance_trx.updated_at, finance_acc.id, finance_acc.name, finance_acc.created_at, finance_acc.updated_at FROM finance_trx JOIN finance_acc ON finance_trx.acc_id = finance_acc.id WHERE finance_trx.id = ? AND finance_trx.is_deleted = 0;",
		vid)

	err := row.Scan(
		&result.ID,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Account.ID,
		&result.Account.Name,
		&result.Account.CreatedAt,
		&result.Account.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return result, x.WrapWithCode(err, x.CodeHTTPNotFound, "GetSQLFinancialTrxByID Not Found")
		// return result,
	} else if err != nil {
		return result, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancialTrxByID")
	}

	return result, nil
}

func (d *Domain) UpdateSQLFinancialTrx(c *gin.Context, v entity.FinancialTrx) (entity.FinancialTrx, error) {
	_, err := DB.Exec("UPDATE finance_trx SET acc_id = ? WHERE id = ?;",
		v.Account.ID,
		v.UpdatedAt,
		// WHERE
		v.ID)
	if err != nil {
		return v, x.WrapWithCode(err, x.CodeSQLUpdate, "UpdateSQLFinancialTrx")
	}

	return v, nil
}

func (d *Domain) DeleteSQLFinancialTrx(c *gin.Context, vid int64) error {
	_, err := DB.Exec("UPDATE finance_trx SET is_deleted = ? , deleted_at = ? WHERE id = ?;",
		true,
		time.Now(),
		// WHERE
		vid)
	if err != nil {
		return x.WrapWithCode(err, x.CodeSQLDelete, "DeleteSQLFinancialTrx")
	}

	return nil
}
