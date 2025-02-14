package database

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.UserName, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logrus.Println("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Println("Error pinging database: %v", err)
	}

	return db, nil

}
