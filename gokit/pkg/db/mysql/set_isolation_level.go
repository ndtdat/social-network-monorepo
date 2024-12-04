package mysql

import (
	"fmt"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/enum"
	"gorm.io/gorm"
)

func setIsolationLevel(db *gorm.DB, level enum.IsolationLevel) error {
	if level == enum.IsolationLevel_NONE {
		level = enum.IsolationLevel_READ_COMMITTED
	}

	return db.Raw(fmt.Sprintf("SET SESSION TRANSACTION ISOLATION LEVEL %s", level)).Error
}
