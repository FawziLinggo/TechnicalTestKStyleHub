package domain

import (
	"database/sql"
	"gorm.io/gorm"
)

type App struct {
	Name          string
	MaxWorkerPool int
}

type Sql struct {
	DB         *gorm.DB
	Connection *sql.DB
}
