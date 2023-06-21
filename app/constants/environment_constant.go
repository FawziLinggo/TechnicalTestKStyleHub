package constants

const (
	EnvironmentDirectory = "../../.env"

	EnvironmentAppName          = "APP_NAME"
	EnvironmentAppRestPort      = "APP_REST_PORT"
	EnvironmentAppMaxWorkerPool = "APP_WORKER_MAX_POOL"

	EnvironmentLogPostgresMode = "LOG_MYSQL_MODE"

	EnvironmentMySqlMigrationDirectory = "MYSQL_MIGRATION_DIRECTORY"
	EnvironmentMySqlMigrationDialect   = "MYSQL_MIGRATION_DIALECT"
	EnvironmentMySqlDBHost             = "MYSQL_DB_HOST"
	EnvironmentMySqlDBUser             = "MYSQL_DB_USER"
	EnvironmentMySqlDBPassword         = "MYSQL_DB_PASSWORD"
	EnvironmentMySqlDBName             = "MYSQL_DB_NAME"
	EnvironmentMySqlDBPort             = "MYSQL_DB_PORT"
	EnvironmentMySqlDBSSLMode          = "MYSQL_DB_SSL_MODE"
	EnvironmentMySqlDBMaxPool          = "MYSQL_DB_MAX_POOL"
	EnvironmentMySqlDBMinPool          = "MYSQL_DB_MIN_POOL"

	ConstantOriginCORS = "ORIGIN_CORS"

	EnvironmentSecretJWT = "SECRET_JWT"
)
