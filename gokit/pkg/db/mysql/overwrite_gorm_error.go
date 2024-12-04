package mysql

import (
	"errors"
	"fmt"
	gomysql "github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrDuplicatedKey = errors.New("duplicated key not allowed")
)

func overwriteGormError(db *gorm.DB) {
	_ = db.Callback().Create().After("gorm:after_create").Register("custom_error", customGormError)
	_ = db.Callback().Update().After("gorm:after_update").Register("custom_error", customGormError)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("custom_error", customGormError)
	_ = db.Callback().Query().After("gorm:after_query").Register("custom_error", customGormError)
	_ = db.Callback().Raw().After("gorm:after_raw").Register("custom_error", customGormError)
	_ = db.Callback().Row().After("gorm:after_row").Register("custom_error", customGormError)
}

func customGormError(db *gorm.DB) {
	if db.Error != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(db.Error, &mysqlErr) && mysqlErr.Number == gomysql.ER_DUP_ENTRY {
			db.Error = ErrDuplicatedKey
		} else if !errors.Is(db.Error, gorm.ErrRecordNotFound) {
			db.Error = fmt.Errorf("something went wrong")
		}
	}
}
