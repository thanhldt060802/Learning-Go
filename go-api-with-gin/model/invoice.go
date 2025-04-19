package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Invoice struct {
	InvoiceId   int64           `json:"invoice_id" gorm:"column:invoice_id;"`
	UserId      int64           `json:"user_id" gorm:"column:user_id;"`
	TotalAmount decimal.Decimal `json:"total_amount" gorm:"column:total_amount;"`
	Status      string          `json:"status" gorm:"column:status;"`
	CreatedAt   time.Time       `json:"created_at" gorm:"column:created_at;"`
}
