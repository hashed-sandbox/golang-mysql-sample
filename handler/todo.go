package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hashed-sandbox/golang-mysql-sample/infra"
	"github.com/hashed-sandbox/golang-mysql-sample/repository"
)

type indexJSON struct {
	ID   uint64  `json:"id"`
	Name *string `json:"name"`
}

// TodoIndex is a handler for GET /api/v1/todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	sqlHandler := infra.NewSQLHandler()
	todos := repository.FindAllTodos(sqlHandler)

	result := make([]indexJSON, len(todos))
	for i, v := range todos {
		result[i].ID = v.ID
		result[i].Name = v.Name
	}

	json.NewEncoder(w).Encode(result)
}
