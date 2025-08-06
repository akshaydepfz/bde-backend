package main

import (
	"net/http"

	"lantorabde.app/database"
	"lantorabde.app/handler"
)

func main() {
	database.ConnectDatabase()

	http.Handle("/users", http.HandlerFunc(handler.Userhandler))
	http.ListenAndServe(":8080", nil)

}
