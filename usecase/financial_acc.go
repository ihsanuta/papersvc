package usecase

import (
	"papersvc/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *Usecase) CreateFinancialAcc(c *gin.Context, v entity.FinancialAccReq) (entity.FinancialAcc, error) {
	var financialacc entity.FinancialAcc
	payload := entity.FinancialAcc{
		Name:      v.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	financialacc, err := u.Domain.CreateSQLFinancialACC(c, payload)
	if err != nil {
		return financialacc, nil
	}

	return financialacc, nil
}

func (u *Usecase) GetFinancialAccByID(c *gin.Context, vid int64) (entity.FinancialAcc, error) {
	return u.Domain.GetSQLFinancialACCByID(c, vid)
}

func (u *Usecase) UpdateFinancialAcc(c *gin.Context, v entity.FinancialAccReq) (entity.FinancialAcc, error) {
	var financialacc entity.FinancialAcc
	payload := entity.FinancialAcc{
		ID:        v.ID,
		Name:      v.Name,
		UpdatedAt: time.Now(),
	}

	financialacc, err := u.Domain.UpdateSQLFinancialACC(c, payload)
	if err != nil {
		return financialacc, nil
	}

	return financialacc, nil
}

func (u *Usecase) DeleteFinancialAcc(c *gin.Context, vid int64) error {
	return u.Domain.DeleteSQLFinancialACC(c, vid)
}
