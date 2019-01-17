package main

import (
	"log"
	"net/http"

	"github.com/hashed-sandbox/golang-mysql-sample/handler"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error)
	}
}

func main() {
	http.HandleFunc("/api/v1/todos", handler.Default(handler.TodoIndex))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
