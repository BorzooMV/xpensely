package models

import "time"

type Expense struct {
	Id          int64      `json:"id"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ExpenseWithUserId struct {
	Expense
	UserId int64 `json:"user_id"`
}
