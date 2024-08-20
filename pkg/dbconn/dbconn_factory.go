package dbconn

import (
	"database/sql"
	"errors"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/codespace-id/codespace-x/pkg/dbconn/mysql"
)

func GetDb(db enum.DbType) (*sql.DB, error) {
	if db == enum.MYSQL {
		return mysql.NewMysqlDB()
	}
	return nil, errors.New("db not supported")
}
