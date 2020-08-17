package entity

import "time"

type FinancialAcc struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FinancialAccReq struct {
	ID   int64  `json:"-"`
	Name string `json:"name"`
}

type FinancialAccParam struct {
	ID     int64  `param:"id" db:"id" form:"id"`
	Name   string `param:"name" db:"name" form:"name"`
	SortBy string `param:"sort" db:"sort" form:"sort"`
	Page   int64  `param:"page" db:"page" form:"page"`
	Limit  int64  `param:"limit" db:"limit" form:"limit"`
}
