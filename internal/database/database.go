package database

import (
	"database/sql"
	"fmt"

	"github.com/SunilKividor/drape/internal/configs"
)

type dbServer struct {
	config configs.DBConfig
}

func NewDbServer(config configs.DBConfig) dbServer {
	return dbServer{
		config: config,
	}
}

func (db *dbServer) StartDatabase() (*sql.DB, error) {
	switch db.config.DBDriver {
	case "sql":
		return startSQLDB(db.config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", db.config.DBDriver)
	}

}

func startSQLDB(config configs.DBConfig) (*sql.DB, error) {
	db, err := sqlDB(config)
	return db, err
}
