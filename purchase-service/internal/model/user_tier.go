package model

import "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"

type UserTier struct {
	UserID    uint64                     `gorm:"primaryKey"`
	Tier      model.SubscriptionPlanTier `gorm:"not null;type:enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM')"`
	CreatedAt uint64
	UpdatedAt uint64
}

const (
	UserTierTableName   = "user_tiers"
	UserTier_USER_ID    = "user_id"
	UserTier_TIER       = "tier"
	UserTier_CREATED_AT = "created_at"
	UserTier_UPDATED_AT = "updated_at"
)
