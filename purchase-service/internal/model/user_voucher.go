package model

import "github.com/ndtdat/social-network-monorepo/purchase-service/pkg/api/go/purchase/model"

type UserVoucher struct {
	ID uint64 `gorm:"primaryKey"`

	UserID                 uint64                `gorm:"not null"`
	VoucherConfigurationID uint64                `gorm:"not null"`
	VoucherCode            string                `gorm:"not null;unique"`
	VoucherConfiguration   *VoucherConfiguration `gorm:"foreignKey:VoucherConfigurationID"`

	ExpiredAt uint64 `gorm:"not null"`

	Status    model.UserVoucherStatus `gorm:"not null;index;type:enum('UVS_ALLOCATED','UVS_USED','UVS_EXPIRED')"`
	CreatedAt uint64
	UpdatedAt uint64
}

const (
	UserVoucherTableName                      = "user_vouchers"
	UserVoucher_PRELOAD_VOUCHER_CONFIGURATION = "VoucherConfiguration"
	UserVoucher_ID                            = "id"
	UserVoucher_USER_ID                       = "user_id"
	UserVoucher_VOUCHER_GROUP_ID              = "voucher_group_id"
	UserVoucher_VOUCHER_CODE                  = "voucher_code"
	UserVoucher_STATUS                        = "status"
	UserVoucher_EXPIRED_AT                    = "expired_at"
	UserVoucher_CREATED_AT                    = "created_at"
	UserVoucher_UPDATED_AT                    = "updated_at"
)
