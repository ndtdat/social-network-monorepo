package model

type UserCampaign struct {
	UserID     uint64 `gorm:"primaryKey"`
	CampaignID uint64 `gorm:"primaryKey"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	UserCampaignTableName    = "user_campaigns"
	UserCampaign_CAMPAIGN_ID = "campaign_id"
	UserCampaign_CREATED_AT  = "created_at"
	UserCampaign_UPDATED_AT  = "updated_at"
)
