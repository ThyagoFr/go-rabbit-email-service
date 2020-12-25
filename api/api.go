package api

import (
	"fmt"
	"net/http"
)

type application struct {
	EmailHistoryService interface{}
}

func (app *application) GetHistory(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Executando 123...")
	_, _ = response.Write([]byte("Iniciando aplicacao..."))
}
