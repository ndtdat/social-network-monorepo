package model

type VoucherPool struct {
	Code string `gorm:"primaryKey;size:256;not null"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	VoucherPoolTableName   = "voucher_pools"
	VoucherPool_CODE       = "code"
	VoucherPool_CREATED_AT = "created_at"
	VoucherPool_UPDATED_AT = "updated_at"
)
