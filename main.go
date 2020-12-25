package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/thyagofr/go-rabbit-email-service/database"

	"github.com/thyagofr/go-rabbit-email-service/stream"

	"github.com/joho/godotenv"
	"github.com/thyagofr/go-rabbit-email-service/api"
	"github.com/thyagofr/go-rabbit-email-service/configs"
)

const (
	PROD string = "prod"
	DEV  string = "dev"
)

func init() {
	environment := *flag.String("env", "dev", "-Ambiente de execucao da aplicacao")
	flag.Parse()
	if environment != PROD && environment != DEV {
		log.Fatal("Ambiente de execucao incorreto. Possibilidades : dev - Desenvolvimento, prod - Producao")
	}
	if err := godotenv.Load(fmt.Sprintf(".env-%s", environment)); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config := *configs.InitConfiguration()
	client := database.InitDatabase(&config.MONGODB)
	defer client.Disconnect(context.Background())
	stream.InitStream(&config.RABBITMQ)
	server := api.InitAPI(&config.MONGODB)
	log.Fatal(http.ListenAndServe(config.SERVER.URL(), server))
}
