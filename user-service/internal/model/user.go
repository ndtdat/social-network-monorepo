package model

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Email    string `gorm:"not null;index"`
	Password string `gorm:"not null"`

	CreatedAt uint64
	UpdatedAt uint64
}

const (
	UserTableName   = "users"
	User_ID         = "id"
	User_EMAIL      = "email"
	User_PASSWORD   = "password"
	User_CREATED_AT = "created_at"
	User_UPDATED_AT = "updated_at"
)
