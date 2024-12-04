package mysql

import (
	"database/sql"
	"fmt"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewMysqlDB() (*sql.DB, error) {

	dbHost := config.Host
	dbUser := config.Username
	dbPass := config.Password
	dbName := config.Database

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
