package mysql

import (
	"database/sql"
	"fmt"
	driver "github.com/go-sql-driver/mysql"
	"github.com/ndtdat/social-network-monorepo/gokit/pkg/config"
	"go.uber.org/zap"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	glogger "gorm.io/gorm/logger"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.App, logger *zap.Logger) (*gorm.DB, error) {
	mysqlCfg := cfg.MySQL
	tracingCfg := cfg.Tracing
	tracingEnabled := tracingCfg.Enabled

	dbName := mysqlCfg.DB
	/*
		- Sets the collation used for client-server interaction on connection. In contrast to charset, collation does
		not issue additional queries. If the specified collation is unavailable on the target server, the connection
		will fail.
	*/
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?collation=utf8mb4_0900_ai_ci&parseTime=True&loc=Local&multiStatements=true&"+
			"interpolateParams=%s",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Host, mysqlCfg.Port, dbName,
		strconv.FormatBool(mysqlCfg.InterpolateParams),
	)

	var (
		db    *gorm.DB
		sqlDB *sql.DB
		err   error
	)

	gormCfg := &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	if mysqlCfg.LogMode != 0 {
		gormCfg.Logger = glogger.Default.LogMode(glogger.LogLevel(mysqlCfg.LogMode))
	}

	if mysqlCfg.PreparedStatement {
		gormCfg.PrepareStmt = true
	}

	createBatchSize := mysqlCfg.CreateBatchSize
	if createBatchSize > 0 {
		gormCfg.CreateBatchSize = int(createBatchSize)
	}

	if driverName := "mysql"; tracingEnabled {
		sqltrace.Register(
			driverName, &driver.MySQLDriver{},
			sqltrace.WithServiceName(fmt.Sprintf("mysql.%s", tracingCfg.ServiceName)),
		)

		sqlDB, err = sqltrace.Open(driverName, dsn)
		if err != nil {
			return nil, fmt.Errorf("cannot open sqltrace mysql due to %v", err)
		}

		db, err = gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), gormCfg)
	} else {
		db, err = gorm.Open(mysql.Open(dsn), gormCfg)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot open gorm mysql due to %v", err)
	}

	if mysqlCfg.OverwriteGormError {
		overwriteGormError(db)
	}

	sqlDB, err = db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(int(mysqlCfg.MaxIdleConnections))
	sqlDB.SetMaxOpenConns(int(mysqlCfg.MaxOpenConnections))
	sqlDB.SetConnMaxLifetime(mysqlCfg.ConnectionMaxLifetime)

	migrationCfg := mysqlCfg.Migration
	if migrationCfg.Enabled {
		err = migrate(logger, sqlDB, dbName, migrationCfg.SourceURL)
		if err != nil {
			return nil, fmt.Errorf("cannot migrate MySQL due to %v", err)
		}
	}

	groupConcatMaxLen := mysqlCfg.GroupConcatMaxLen
	if groupConcatMaxLen > 0 {
		err = setGroupConcatMaxLen(db, groupConcatMaxLen)
		if err != nil {
			return nil, fmt.Errorf("cannot set group concat max len due to %v", err)
		}
	}

	return db, setIsolationLevel(db, mysqlCfg.IsolationLevel)
}
