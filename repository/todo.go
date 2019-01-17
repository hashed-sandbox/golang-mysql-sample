package repository

import (
	"github.com/hashed-sandbox/golang-mysql-sample/domain/rdb"
	"github.com/hashed-sandbox/golang-mysql-sample/infra"
)

func FindAllTodos(sqlHandler *infra.SQLHandler) (todos []rdb.Todo) {
	sess := sqlHandler.Conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		panic(err.Error)
	}

	tx.Select("*").From("todos").Load(&todos)

	return
}
