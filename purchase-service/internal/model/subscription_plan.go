package model

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"github.com/shopspring/decimal"
)

type SubscriptionPlan struct {
	Tier           model.SubscriptionPlanTier `gorm:"primaryKey;type:enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM')"`
	CurrencySymbol string                     `gorm:"index;not null"`
	Price          decimal.Decimal            `gorm:"type:decimal(23,8)"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	SubscriptionPlanTableName        = "subscription_plans"
	SubscriptionPlan_TIER            = "tier"
	SubscriptionPlan_CURRENCY_SYMBOL = "currency_symbol"
	SubscriptionPlan_PRICE           = "price"
	SubscriptionPlan_CREATED_AT      = "created_at"
	SubscriptionPlan_UPDATED_AT      = "updated_at"
)
