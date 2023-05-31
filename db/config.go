package db

import (
	"Labora-Wallet/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	var err error

	DbConfig, err := models.LoadEnvVariables()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	DbConfig.Host, DbConfig.Port, DbConfig.User, DbConfig.Password, DbConfig.DbName)

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		DbConfig.User, DbConfig.Password, DbConfig.Host, DbConfig.Port, DbConfig.DbName)

	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return connection, err
}
