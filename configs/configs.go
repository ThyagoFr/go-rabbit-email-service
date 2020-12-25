package configs

import (
	"fmt"
	"os"
)

type ApplicationConfig struct {
	SERVER   ServerConfig
	RABBITMQ RabbitMQConfig
	MONGODB  MongoDBConfig
}

type ServerConfig struct {
	HOST string
	PORT string
}

func (sr *ServerConfig) URL() string {
	url := fmt.Sprintf("%s:%s", sr.HOST, sr.PORT)
	return url
}

type RabbitMQConfig struct {
	HOST     string
	PORT     string
	USER     string
	PASSWORD string
	QUEUE    string
	CONSUMER string
	DLXQUEUE string
	DLXTIME  string
}

func (rb *RabbitMQConfig) URL() (url string) {
	url = fmt.Sprintf("amqp://%s:%s@%s:%s/", rb.USER, rb.PASSWORD, rb.HOST, rb.PORT)
	return
}

type MongoDBConfig struct {
	HOST           string
	PORT           string
	USERNAME       string
	PASSWORD       string
	DABABASENAME   string
	COLLECTIONNAME string
}

func (db *MongoDBConfig) URL() (url string) {
	url = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", db.USERNAME, db.PASSWORD, db.HOST, db.PORT, db.DABABASENAME)
	return
}

var configuration *ApplicationConfig

func InitConfiguration() *ApplicationConfig {
	configuration = &ApplicationConfig{
		SERVER: ServerConfig{
			HOST: os.Getenv("SERVER_HOST"),
			PORT: os.Getenv("SERVER_PORT"),
		},
		RABBITMQ: RabbitMQConfig{
			HOST:     os.Getenv("RABBITMQ_HOST"),
			PORT:     os.Getenv("RABBITMQ_PORT"),
			USER:     os.Getenv("RABBITMQ_USER"),
			PASSWORD: os.Getenv("RABBITMQ_PASSWORD"),
			QUEUE:    os.Getenv("RABBITMQ_QUEUE"),
			CONSUMER: os.Getenv("RABBITMQ_CONSUMER"),
			DLXQUEUE: os.Getenv("RABBITMQ_DLXQUEUE"),
			DLXTIME:  os.Getenv("RABBITMQ_DLXTIME"),
		},
		MONGODB: MongoDBConfig{
			HOST:           os.Getenv("MONGODB_HOST"),
			PORT:           os.Getenv("MONGODB_PORT"),
			USERNAME:       os.Getenv("MONGODB_USERNAME"),
			PASSWORD:       os.Getenv("MONGODB_PASSWORD"),
			DABABASENAME:   os.Getenv("MONGODB_DATABASENAME"),
			COLLECTIONNAME: os.Getenv("MONGODB_COLLECTIONNAME"),
		},
	}
	return configuration
}
