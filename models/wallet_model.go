package models

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Wallet struct {
	ID        int    `json:"id"`
	CountryId string `json:"countryid"`
	DNI       int    `json:"dni"`
	Date      string `json:"date"`
	Balance   int    `json:"balance"`
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func LoadEnvVariables() (DbConfig, error) {
	err := godotenv.Load(".env") //metodo para cargar nuestras variables de un file.
	if err != nil {
		log.Fatalf("Error loading .env file")
		return DbConfig{}, err
	}
	return DbConfig{
		Host:     os.Getenv("host"),
		Port:     os.Getenv("port"),
		User:     os.Getenv("postgresql_user"),
		Password: os.Getenv("postgresql_password"),
		DbName:   os.Getenv("db_name"),
	}, nil
}

func LoadEnvVariablesTruora() (TruoraConfig, error) {
	err := godotenv.Load(".env") //metodo para cargar nuestras variables de un file.
	if err != nil {
		log.Fatalf("Error loading .env file")
		return TruoraConfig{}, err
	}
	return TruoraConfig{
		TruoraApiKey: os.Getenv("truora_api_key"),
		Urlbase:      os.Getenv("urlbasetruora"),
	}, nil
}

type TruoraConfig struct {
	TruoraApiKey string
	Urlbase      string
}

type Transactions struct {
	SenderId   int    `json:"senderid"`
	ReceiverId int    `json:"receiverid"`
	Amount     int    `json:"amount"`
	Operation  string `json:"operation"`
	Date       string `json:"date"`
}

type Movements struct {
	Operation string `json:"operation"`
	Amount    int    `json:"amount"`
	Time      string `json:"date"`
}

type TransactionInfo struct {
	ID        int         `json:"id"`
	Balance   int         `json:"balance"`
	Movements []Movements `json:"movements"`
}
