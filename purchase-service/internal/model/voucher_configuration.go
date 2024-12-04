package model

import (
	"github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"
	"github.com/shopspring/decimal"
)

type VoucherConfiguration struct {
	ID uint64 `gorm:"primaryKey"`

	CalculationType model.VoucherCalculationType `gorm:"not null;index;type:enum('VCT_PERCENTAGE','VCT_AMOUNT')"`
	CurrencySymbol  string                       `gorm:"default:null"`
	Value           decimal.Decimal              `gorm:"type:decimal(23,8)"`

	MaxQty       uint32 `gorm:" not null"`
	AllocatedQty uint32 `gorm:" not null"`
	RedeemedQty  uint32 `gorm:" not null"`

	StartAt uint64 `gorm:"not null"`
	EndAt   uint64 `gorm:"not null"`

	Status model.VoucherStatus `gorm:"not null;type:enum('VS_DRAFT','VS_AVAILABLE','VS_UNAVAILABLE')"`

	AppliedTier model.SubscriptionPlanTier `gorm:"not null;type:enum('SPT_BRONZE','SPT_SILVER','SPT_GOLD','SPT_PLATINUM')"`

	Type       model.VoucherGroupType `gorm:"not null;index;type:enum('VGT_CAMPAIGN')"`
	CampaignID uint64                 `gorm:"default:null"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	VoucherConfigurationTableName         = "voucher_configurations"
	VoucherConfiguration_ID               = "id"
	VoucherConfiguration_CALCULATION_TYPE = "calculation_type"
	VoucherConfiguration_CURRENCY_SYMBOL  = "currency_symbol"
	VoucherConfiguration_VALUE            = "value"
	VoucherConfiguration_MAX_QTY          = "max_qty"
	VoucherConfiguration_ALLOCATED_QTY    = "allocated_qty"
	VoucherConfiguration_REDEEMED_QTY     = "redeemed_qty"
	VoucherConfiguration_START_AT         = "start_at"
	VoucherConfiguration_END_AT           = "end_at"
	VoucherConfiguration_STATUS           = "status"
	VoucherConfiguration_TYPE             = "type"
	VoucherConfiguration_CAMPAIGN_ID      = "campaign_id"
	VoucherConfiguration_CREATED_AT       = "created_at"
	VoucherConfiguration_UPDATED_AT       = "updated_at"
)
