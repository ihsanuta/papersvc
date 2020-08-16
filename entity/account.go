package entity

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type ReqAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
