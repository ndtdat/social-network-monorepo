package mysql

import (
	"database/sql"
	"fmt"
	gmigrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	//nolint:revive
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
)

func migrate(logger *zap.Logger, db *sql.DB, dbName, sourceURL string) error {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := gmigrate.NewWithDatabaseInstance(
		sourceURL,
		dbName,
		driver,
	)
	if err != nil {
		return err
	}

	//nolint:nestif
	logger.Info("MySQL migration is starting")
	if err = m.Up(); err != nil {
		switch err {
		case gmigrate.ErrNoChange:
			logger.Info("No new changes to apply to the database")
		default:
			return fmt.Errorf("db migration fails due to %v ", err)
		}
	}

	logger.Info("MySQL is migrated successfully")

	return nil
}
