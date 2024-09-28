package database

import (
	"database/sql"
	"fmt"

	"github.com/SunilKividor/drape/internal/configs"
	_ "github.com/lib/pq"
)

func sqlDB(config configs.DBConfig) (*sql.DB, error) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	return db, err
}
