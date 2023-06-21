package libraries

import (
	"TechnicalTestKStyleHub/app/constants"
	"database/sql"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"time"
)

type GormLibrary struct {
	MigrationDirectory string
	MigrationDialect   string
	DBHost             string
	DBUser             string
	DBPassword         string
	DBPort             string
	DBName             string
	DBSSLMode          string
	MaxPool            int
	MinPool            int
	LogMode            int
}

func (lib GormLibrary) Migrate(db *sql.DB) (err error) {

	// init
	currentWorkDirectory, _ := os.Getwd()
	migrations := &migrate.FileMigrationSource{
		Dir: fmt.Sprintf("%s/%s", currentWorkDirectory, lib.MigrationDirectory),
	}

	// run
	n, err := migrate.Exec(db, lib.MigrationDialect, migrations, migrate.Up)
	if err != nil {
		return err
	}

	// res
	if n > constants.EmptyNumber {
		fmt.Println(fmt.Sprintf("Migration: Applied %d migrations", n))
	} else {
		fmt.Println("Migration: No schema change")
	}

	return err
}
func (lib GormLibrary) ConnectAndValidate() (db *gorm.DB, sql *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		lib.DBUser,
		lib.DBPassword,
		lib.DBHost,
		lib.DBPort,
		lib.DBName,
	)
	// log mode
	logConfig := logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
	}
	switch lib.LogMode {
	case 1:
		logConfig.LogLevel = logger.Info
		logConfig.IgnoreRecordNotFoundError = false
	case 2:
		logConfig.LogLevel = logger.Warn
		logConfig.IgnoreRecordNotFoundError = true
	case 3:
		logConfig.LogLevel = logger.Error
		logConfig.IgnoreRecordNotFoundError = true
	default:
		logConfig = logger.Config{}
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(lib.LogMode)),
	}

	db, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	sqlDB.SetMaxIdleConns(lib.MinPool)
	sqlDB.SetMaxOpenConns(lib.MaxPool)

	err = sqlDB.Ping()
	if err != nil {
		return nil, nil, err
	}

	return db, sqlDB, err
}
