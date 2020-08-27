package main

import (
	"angular-code-share/api/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/create", handlers.Create)

	http.HandleFunc("/get_by_id", handlers.GetByID)

	http.HandleFunc("/update_by_id", handlers.UpdateByID)

	http.ListenAndServe(":8081", nil)

}
