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
