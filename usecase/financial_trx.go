package usecase

import (
	"time"

	"github.com/ihsanuta/papersvc/entity"

	x "github.com/ihsanuta/papersvc/errors"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) CreateFinancialTrx(c *gin.Context, v entity.FinancialTrxReq) (entity.FinancialTrx, error) {
	var financialtrx entity.FinancialTrx

	// Cek Account ID Exist
	acc, err := u.Domain.GetSQLFinancialACCByID(c, v.AccountID)
	if err != nil && x.ErrCode(err) != x.CodeHTTPNotFound {
		return financialtrx, x.Wrap(err, "ErrorGetSQLFinancialACCByID")
	}

	if acc.ID < 1 {
		return financialtrx, x.NewWithCode(x.CodeHTTPBadRequest, "ErrorGetSQLFinancialACCByID")
	}

	payload := entity.FinancialTrx{
		Account: entity.FinancialAcc{
			ID: v.AccountID,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	financialtrx, err = u.Domain.CreateSQLFinancialTrx(c, payload)
	if err != nil {
		return financialtrx, nil
	}

	financialtrx.Account, err = u.Domain.GetSQLFinancialACCByID(c, payload.Account.ID)
	if err != nil {
		return financialtrx, nil
	}

	return financialtrx, nil
}

func (u *Usecase) GetFinancialTrxByID(c *gin.Context, vid int64) (entity.FinancialTrx, error) {
	return u.Domain.GetSQLFinancialTrxByID(c, vid)
}

func (u *Usecase) UpdateFinancialTrx(c *gin.Context, v entity.FinancialTrxReq) (entity.FinancialTrx, error) {
	var financialtrx entity.FinancialTrx

	// Cek Account ID Exist
	acc, err := u.Domain.GetSQLFinancialACCByID(c, v.AccountID)
	if err != nil && x.ErrCode(err) != x.CodeHTTPNotFound {
		return financialtrx, x.Wrap(err, "ErrorGetSQLFinancialACCByID")
	}

	if acc.ID < 1 {
		return financialtrx, x.NewWithCode(x.CodeHTTPBadRequest, "ErrorGetSQLFinancialACCByID")
	}

	trx, err := u.Domain.GetSQLFinancialTrxByID(c, v.ID)
	if err != nil && x.ErrCode(err) != x.CodeHTTPNotFound {
		return financialtrx, x.Wrap(err, "GetSQLFinancialTrxByID")
	}

	if trx.ID < 1 {
		return financialtrx, x.NewWithCode(x.CodeHTTPBadRequest, "ErrorGetSQLFinancialTrxByID")
	}

	payload := entity.FinancialTrx{
		ID: v.ID,
		Account: entity.FinancialAcc{
			ID: v.AccountID,
		},
		CreatedAt: trx.CreatedAt,
		UpdatedAt: time.Now(),
	}

	financialtrx, err = u.Domain.UpdateSQLFinancialTrx(c, payload)
	if err != nil {
		return financialtrx, nil
	}

	financialtrx.Account, err = u.Domain.GetSQLFinancialACCByID(c, payload.Account.ID)
	if err != nil {
		return financialtrx, nil
	}

	return financialtrx, nil
}

func (u *Usecase) DeleteFinancialTrx(c *gin.Context, vid int64) error {
	return u.Domain.DeleteSQLFinancialTrx(c, vid)
}

func (u *Usecase) GetFinancialTrx(c *gin.Context, param entity.FinancialTrxParam) ([]entity.FinancialTrx, entity.Pagination, error) {
	return u.Domain.GetSQLFinancialTrx(c, param)
}
