package main

import (
	"Sclad/api"
	"Sclad/db"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// NewHTTPHandler Функция создания экземпляра структуры (единственного!)
func NewHTTPHandler(database db.DataBaseEvents) *api.Handlers {
	return &api.Handlers{Database: database}
}

// NewServer Функция создания сервера
func NewServer() *http.Server {

	//создаем структуры бд (для inmemory инициализируем map сразу)
	var mon, mem db.DataBaseEvents = &db.Mongo{}, &db.Memory{
		Posts:   make(map[string]db.Post),
		Authors: make(map[string][]string),
	}

	//создаем структуру, в которой будут реализованы маршруты (и хранится ссылка на бд)
	var handler *api.Handlers

	//получаем выбранное хранилище
	flag := "inmemory"

	//если флаг "mongo" то стартуем монгу, если "inmemory" - стартуем локальное
	if flag == "mongo" {
		mon.InitDB()
		handler = NewHTTPHandler(mon)

	} else if flag == "inmemory" {
		mem.InitDB()
		handler = NewHTTPHandler(mem)

	}

	//создаем роутер (Gorilla/mux)
	r := mux.NewRouter()

	//добавляем маршруты
	r.HandleFunc("/api/v1/add", handler.AddItem).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/get/item/{itemId:[A-Za-z0-9_\\-]+}", handler.GetItem).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/get/{userId}", handler.GetItems).Methods(http.MethodGet,"OPTIONS")

	//создаем http соединение
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("0.0.0.0:%s", "8000"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv
}

func main() {
	srv := NewServer()
	log.Printf("Start serving on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
