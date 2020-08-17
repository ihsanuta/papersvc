package entity

import "time"

type FinancialTrx struct {
	ID        int64        `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Account   FinancialAcc `json:"account"`
}

type FinancialTrxReq struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
}

type FinancialTrxParam struct {
	ID        int64  `param:"id" db:"finance_trx.id" form:"id"`
	AccountID int64  `param:"account_id" db:"finance_trx.acc_id" form:"account_id"`
	SortBy    string `param:"sort" db:"sort" form:"sort"`
	Page      int64  `param:"page" db:"page" form:"page"`
	Limit     int64  `param:"limit" db:"limit" form:"limit"`
}
