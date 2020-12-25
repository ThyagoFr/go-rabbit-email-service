package api

import (
	"github.com/gorilla/mux"
	"github.com/thyagofr/go-rabbit-email-service/configs"
)

func InitAPI(dbConfig *configs.MongoDBConfig) *mux.Router {
	app := application{
		EmailHistoryService: "Servico ficticio",
	}
	router := mux.NewRouter().StrictSlash(true)
	api := router.PathPrefix("/api/v1/emails").Subrouter()
	api.HandleFunc("", app.GetHistory).Methods("GET")
	return router
}
