package infra

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const MAX_OPEN_CONNS = 10

type SQLHandler struct {
	Conn *dbr.Connection
}

func getDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%%2FTokyo",
		os.Getenv("RDB_USER"),
		os.Getenv("RDB_PASSWORD"),
		os.Getenv("RDB_HOST"),
		os.Getenv("RDB_PORT"),
		os.Getenv("RDB_DATABASE"),
	)
}

func NewSQLHandler() *SQLHandler {
	dataSourceName := getDataSourceName()
	conn, err := dbr.Open("mysql", dataSourceName, nil)
	if err != nil {
		panic(err.Error)
	}

	conn.SetMaxOpenConns(MAX_OPEN_CONNS)

	return &SQLHandler{Conn: conn}
}
