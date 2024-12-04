package dbconn

import (
	"database/sql"
	"errors"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/enum"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/dbconn/mysql"
)

func GetDb(db enum.DbType) (*sql.DB, error) {
	if db == enum.MYSQL {
		return mysql.NewMysqlDB()
	}
	return nil, errors.New("db not supported")
}
