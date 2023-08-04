package mysql

import (
	"database/sql"
	"log"
	cf "xsis/movie/common/config"

	"github.com/go-sql-driver/mysql"
)

func NewMysql(cfg *cf.Config) *sql.DB {
	conn := mysql.Config{
		User:                 cfg.MySQL.User,
		Passwd:               cfg.MySQL.Passwd,
		Net:                  cfg.MySQL.Net,
		Addr:                 cfg.MySQL.Addr,
		DBName:               cfg.MySQL.DBName,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", conn.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	return db
}
