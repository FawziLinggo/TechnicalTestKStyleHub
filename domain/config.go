package domain

import (
	"TechnicalTestKStyleHub/app/constants"
	"TechnicalTestKStyleHub/app/libraries"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	App App
	Sql Sql
}

func LoadConfiguration() (config Config, err error) {

	// load env
	err = godotenv.Load(constants.EnvironmentDirectory)
	if err != nil {
		return config, err
	}

	// logs formatter
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// app
	maxWorker, err := strconv.Atoi(os.Getenv(constants.EnvironmentAppMaxWorkerPool))
	if err != nil {
		return config, err
	}
	config.App = App{
		Name:          os.Getenv(constants.EnvironmentAppName),
		MaxWorkerPool: maxWorker,
	}

	// postgres
	dbLogMode, err := strconv.Atoi(os.Getenv(constants.EnvironmentLogPostgresMode))
	if err != nil {
		return config, err
	}
	maxPool, err := strconv.Atoi(os.Getenv(constants.EnvironmentMySqlDBMaxPool))
	if err != nil {
		return config, err
	}
	minPool, err := strconv.Atoi(os.Getenv(constants.EnvironmentMySqlDBMinPool))
	if err != nil {
		return config, err
	}
	gormLibrary := libraries.GormLibrary{
		MigrationDirectory: os.Getenv(constants.EnvironmentMySqlMigrationDirectory),
		MigrationDialect:   os.Getenv(constants.EnvironmentMySqlMigrationDialect),
		DBHost:             os.Getenv(constants.EnvironmentMySqlDBHost),
		DBUser:             os.Getenv(constants.EnvironmentMySqlDBUser),
		DBPassword:         os.Getenv(constants.EnvironmentMySqlDBPassword),
		DBPort:             os.Getenv(constants.EnvironmentMySqlDBPort),
		DBName:             os.Getenv(constants.EnvironmentMySqlDBName),
		DBSSLMode:          os.Getenv(constants.EnvironmentMySqlDBSSLMode),
		MaxPool:            maxPool,
		MinPool:            minPool,
		LogMode:            dbLogMode,
	}
	config.Sql.DB, config.Sql.Connection, err = gormLibrary.ConnectAndValidate()
	if err != nil {
		return config, err
	}
	if err = gormLibrary.Migrate(config.Sql.Connection); err != nil {
		return config, err
	}

	return config, err
}
