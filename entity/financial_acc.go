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
	ID     int64    `param:"id" db:"id"`
	Name   int64    `param:"name" db:"name"`
	SortBy []string `param:"sort_by" db:"sort_by"`
	Page   int64    `param:"page" db:"page"`
	Limit  int64    `param:"limit" db:"limit"`
}
