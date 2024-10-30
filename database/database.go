package database

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/config"
)

func NewRepository(cfg *config.Config) (*sql.DB, error) {
	sqlCfg := mysql.Config{User: cfg.DBUsername, Passwd: cfg.DBPassword, DBName: cfg.DBName}
	db, err := sql.Open("mysql", sqlCfg.FormatDSN()); if  err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id VARCHAR(64) PRIMARY KEY,
		username VARCHAR(16) NOT NULL,
        email VARCHAR(64) UNIQUE NOT NULL,
        password VARCHAR(64) NOT NULL,
        display_name VARCHAR(128) NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`); if err != nil {
		return nil, err
	}
	
    return db, nil
}