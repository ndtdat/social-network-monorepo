package model

import (
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID             uint64 `gorm:"primaryKey"`
	UserID         uint64 `gorm:"not null;index"`
	CurrencySymbol string `gorm:"not null"`

	OriginalAmount decimal.Decimal `gorm:"not null;type:decimal(23,8')"`
	DiscountAmount decimal.Decimal `gorm:"not null;type:decimal(23,8')"`
	ActualAmount   decimal.Decimal `gorm:"not null;type:decimal(23,8')"`

	SubscriptionPlanID uint64 `gorm:"not null"`
	UserVoucherID      uint64 `gorm:"default null"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	TransactionTableName             = "transactions"
	Transaction_ID                   = "id"
	Transaction_USER_ID              = "user_id"
	Transaction_CURRENCY_SYMBOL      = "currency_symbol"
	Transaction_ORIGINAL_AMOUNT      = "original_amount"
	Transaction_ACTUAL_AMOUNT        = "actual_amount"
	Transaction_SUBSCRIPTION_PLAN_ID = "subscription_plan_id"
	Transaction_USER_VOUCHER_ID      = "user_voucher_id"
	Transaction_CREATED_AT           = "created_at"
	Transaction_UPDATED_AT           = "updated_at"
)
