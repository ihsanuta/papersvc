package domain

import (
	"database/sql"
	"log"
	"time"

	"github.com/ihsanuta/papersvc/entity"
	x "github.com/ihsanuta/papersvc/errors"

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

func (d *Domain) GetSQLFinancialTrx(c *gin.Context, param entity.FinancialTrxParam) ([]entity.FinancialTrx, entity.Pagination, error) {
	results := []entity.FinancialTrx{}

	// validate limit
	if param.Limit < 1 {
		param.Limit = 10
	}

	// validate page
	if param.Page < 1 {
		param.Page = 1
	}

	// set default pagination
	pagination := entity.Pagination{
		CurrentPage:     param.Page,
		CurrentElements: 0,
		TotalPages:      0,
		TotalElements:   0,
		SortBy:          "",
	}

	query, queryc, args, sort := qBuilder(param, "finance_acc.is_deleted=0 AND finance_trx.is_deleted=0")
	rows, err := DB.Query("SELECT finance_trx.id, finance_trx.created_at, finance_trx.updated_at, finance_acc.id, finance_acc.name, finance_acc.created_at, finance_acc.updated_at FROM finance_trx JOIN finance_acc ON finance_trx.acc_id = finance_acc.id "+query, args...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, x.CodeSQLRead, "GetSQLFinancialTrx")
	}
	defer rows.Close()

	for rows.Next() {
		var result entity.FinancialTrx
		if err := rows.Scan(
			&result.ID,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.Account.ID,
			&result.Account.Name,
			&result.Account.CreatedAt,
			&result.Account.UpdatedAt); err != nil {
			return results, pagination, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancialTrx")
		}
		results = append(results, result)
	}

	// Get Invitation Counts
	var totalRecords int64
	row := DB.QueryRow("SELECT COUNT(*) AS total FROM finance_trx JOIN finance_acc ON finance_trx.acc_id = finance_acc.id "+queryc, args...)
	err = row.Scan(
		&totalRecords,
	)

	if err != nil {
		return results, pagination, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancialTrx")
	}

	// Update Pagination
	totalPage := totalRecords / param.Limit
	if totalRecords%param.Limit > 0 || totalRecords == 0 {
		totalPage++
	}
	pagination.TotalPages = totalPage
	pagination.CurrentElements = int64(len(results))
	pagination.TotalElements = totalRecords
	pagination.SortBy = sort

	return results, pagination, nil
}
