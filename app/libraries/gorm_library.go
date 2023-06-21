package libraries

import (
	"TechnicalTestKStyleHub/app/constants"
	"database/sql"
	"errors"
	"fmt"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
type mysqlLogger struct{}

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
		LogLevel:      logger.Info,
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

	//config := &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.LogLevel(lib.LogMode)),
	//}
	gormConfig := gorm.Config{Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logConfig)}

	db, err = gorm.Open(mysql.Open(dsn), &gormConfig)
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

func (lib GormLibrary) InitMigrationFromFile(db *gorm.DB, filePath string) error {
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	var count int64
	result := db.Table("member").Count(&count)
	if result.Error != nil {
		return result.Error
	}
	if count == 0 {
		sqlCommands := strings.Split(string(fileContent), ";")

		for _, cmd := range sqlCommands {
			sql := strings.TrimSpace(cmd)

			if sql != "" {
				result := db.Exec(sql)

				if result.Error != nil {
					return result.Error
				}

				if result.RowsAffected == 0 {
					return errors.New("failed to execute SQL command: " + sql)
				}
			}
		}
	}

	return nil
}
