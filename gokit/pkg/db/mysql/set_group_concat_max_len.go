package mysql

import (
	"fmt"
	"gorm.io/gorm"
)

func setGroupConcatMaxLen(db *gorm.DB, value uint32) error {
	return db.Raw(fmt.Sprintf("SET SESSION group_concat_max_len = %v", value)).Error
}
