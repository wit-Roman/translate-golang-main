package main

import (
	"log"
	"net/http"
	"runtime"

	"translate/constants"
	"translate/db"
	"translate/router"
)

func init() {
	db.Connect()
	db.CreateTables()
}

func main() {
	runtime.GOMAXPROCS(1)

	defer db.Connection.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/orders", router.Orders)

	log.Println("Запуск сервера " + constants.Port)
	err := http.ListenAndServe(constants.Port, mux)

	log.Fatal(err)
}
