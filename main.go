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

	mux.HandleFunc("/session_create", router.SessionCreate)

	mux.HandleFunc("/orders_data", router.OrdersData)
	mux.HandleFunc("/orders_create", router.OrdersCreate)
	mux.HandleFunc("/orders_getById", router.OrderGetById)
	mux.HandleFunc("/orders_update", router.OrdersUpdate)
	mux.HandleFunc("/orders_delete", router.OrdersDelete)
	mux.HandleFunc("/orders_cellEdit", router.OrdersCellEdit)
	mux.HandleFunc("/orders_listen", router.ListenHandler)

	log.Println("Запуск сервера " + constants.Port)
	err := http.ListenAndServe(constants.Port, mux)

	log.Fatal(err)
}
