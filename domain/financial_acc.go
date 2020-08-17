package domain

import (
	"database/sql"
	"fmt"
	"papersvc/entity"
	"time"

	"github.com/gin-gonic/gin"

	x "papersvc/errors"
)

func (d *Domain) CreateSQLFinancialACC(c *gin.Context, v entity.FinancialAcc) (entity.FinancialAcc, error) {
	row, err := DB.Exec("INSERT INTO finance_acc (name, created_at, updated_at) VALUES (?,?,?);",
		v.Name,
		v.CreatedAt,
		v.UpdatedAt)
	if err != nil {
		return v, x.WrapWithCode(err, x.CodeSQLCreate, "CreateSQLFinancialACC")
	}

	v.ID, err = row.LastInsertId()
	if err != nil {
		return v, x.WrapWithCode(err, x.CodeSQLCannotRetrieveLastInsertID, `ErrorGetFinancialACCID`)
	}

	return v, nil
}

func (d *Domain) GetSQLFinancialACCByID(c *gin.Context, vid int64) (entity.FinancialAcc, error) {
	var result entity.FinancialAcc
	row := DB.QueryRow("SELECT id, name, created_at, updated_at FROM finance_acc WHERE id = ? AND is_deleted = 0;",
		vid)

	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return result, x.WrapWithCode(err, x.CodeHTTPNotFound, "GetSQLFinancialACCByID Not Found")
		// return result,
	} else if err != nil {
		return result, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancialACCByID")
	}

	return result, nil
}

func (d *Domain) UpdateSQLFinancialACC(c *gin.Context, v entity.FinancialAcc) (entity.FinancialAcc, error) {
	_, err := DB.Exec("UPDATE finance_acc SET name = ? , updated_at = ? WHERE id = ?;",
		v.Name,
		v.UpdatedAt,
		// WHERE
		v.ID)
	if err != nil {
		return v, x.WrapWithCode(err, x.CodeSQLUpdate, "UpdateSQLFinancialACC")
	}

	return v, nil
}

func (d *Domain) DeleteSQLFinancialACC(c *gin.Context, vid int64) error {
	_, err := DB.Exec("UPDATE finance_acc SET is_deleted = ? , deleted_at = ? WHERE id = ?;",
		true,
		time.Now(),
		// WHERE
		vid)
	if err != nil {
		return x.WrapWithCode(err, x.CodeSQLDelete, "DeleteSQLFinancialACC")
	}

	return nil
}

func (d *Domain) GetSQLFinancial(c *gin.Context, param entity.FinancialAccParam) ([]entity.FinancialAcc, entity.Pagination, error) {
	results := []entity.FinancialAcc{}

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

	query, queryc, args, sort := qBuilder(param, "is_deleted=0")
	fmt.Println("SELECT id, name, created_at, updated_at FROM finance_acc " + query)
	rows, err := DB.Query("SELECT id, name, created_at, updated_at FROM finance_acc "+query, args...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, x.CodeSQLRead, "GetSQLFinancial")
	}
	defer rows.Close()

	for rows.Next() {
		var result entity.FinancialAcc
		if err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.CreatedAt,
			&result.UpdatedAt); err != nil {
			return results, pagination, x.WrapWithCode(err, x.CodeSQLRowScan, "GetSQLFinancial")
		}
		results = append(results, result)
	}

	// Get Invitation Counts
	var totalRecords int64
	row := DB.QueryRow("SELECT COUNT(*) AS total FROM finance_acc "+queryc, args...)
	err = row.Scan(
		&totalRecords,
	)

	if err != nil {
		return results, pagination, x.WrapWithCode(err, x.CodeSQLRowScan, "GetCountSQLFinancial")
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
