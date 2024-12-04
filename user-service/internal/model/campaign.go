package model

import pbmodel "github.com/ndtdat/social-network-monorepo/user-service/pkg/api/go/user/model"

type Campaign struct {
	ID     uint64                 `gorm:"primaryKey"`
	Code   string                 `gorm:"not null"`
	Status pbmodel.CampaignStatus `gorm:"not null;index;type:enum('CS_AVAILABLE','CS_UNAVAILABLE')"`

	StartAt uint64 `gorm:"not null"`
	EndAt   uint64 `gorm:"not null"`

	MaxQty    uint32 `gorm:"not null"`
	JoinedQty uint32 `gorm:"not null"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	CampaignTableName   = "campaigns"
	Campaign_ID         = "id"
	Campaign_CODE       = "code"
	Campaign_STATUS     = "status"
	Campaign_START_AT   = "start_at"
	Campaign_END_AT     = "end_at"
	Campaign_MAX_QTY    = "max_qty"
	Campaign_JOINED_QTY = "joined_qty"
	Campaign_CREATED_AT = "created_at"
	Campaign_UPDATED_AT = "updated_at"
)
