package main

import (
	"net/http"

	"lantorabde.app/database"
	"lantorabde.app/handler"
)

func main() {
	database.ConnectDatabase()

	http.Handle("/users", http.HandlerFunc(handler.Userhandler))
	http.Handle("/users/get", http.HandlerFunc(handler.Userhandlerget))
	http.Handle("/login", http.HandlerFunc(handler.Loginhandler))
	
	http.ListenAndServe(":8080", nil)

}
