package dto

import (
	"github.com/shopspring/decimal"
)

type CreateInvoiceDTO struct {
	UserId      int64           `json:"user_id"`
	TotalAmount decimal.Decimal `json:"total_amount"`
	Status      string          `json:"status"`
}

type UpdateInvoiceDTO struct {
	Status string `json:"status"`
}
