package repositories

import (
	"database/sql"
	"fmt"

	"github.com/SunilKividor/drape/internal/configs"
	_ "github.com/lib/pq"
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
	case "PostgreSQL":
		return startPostgresSQLDB(db.config)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", db.config.DBDriver)
	}

}

func startPostgresSQLDB(config configs.DBConfig) (*sql.DB, error) {
	db, err := postgresSQLDB(config)
	return db, err
}

//establishing connections

func postgresSQLDB(config configs.DBConfig) (*sql.DB, error) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	return db, err
}
