package services

import (
	"Labora-Wallet/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const truoraApiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k1Yzk5MzdiMzBjYzk1YmYxMWQyZmQxN2VkODgxMDkxOCIsImV4cCI6MzI2MTU5NDcxMCwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ3OTQ3MTAsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX1djMFFoU2hyayIsImp0aSI6IjVkNjE3YmRmLTY0ZjItNGMwNC1iZGQzLWVlNWQ4ZTRkYWJmMyIsImtleV9uYW1lIjoicGF1bGEiLCJrZXlfdHlwZSI6ImJhY2tlbmQiLCJ1c2VybmFtZSI6ImdtYWlscGF1MjQ5NC1wYXVsYSJ9.hS2JWBOLyMmMOr5zO4UGklJdfSmk0LuF0TNUEjQ6jR8"

const urlBase = "https://api.checks.truora.com/v1/checks/"

type API struct {
}

func (api *API) ObtainCheckID(dni int, countryId string) (string, error) {
	method := "POST"

	datos := fmt.Sprintf("national_id=%d&country=%s&type=person&user_authorized=true", dni, countryId)
	payload := strings.NewReader(datos)

	client := &http.Client{}
	req, err := http.NewRequest(method, urlBase, payload)

	req.Header.Add("Truora-API-Key", truoraApiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	var backgroundCheck models.BackgroundCheck
	err = json.NewDecoder(res.Body).Decode(&backgroundCheck)
	checkID := backgroundCheck.Check.CheckID

	return checkID, err
}

// falta ObtainScore
// que haga el CreateWallet dependiendo del score
// en mi package db (insert) hacer los logs de los intentos (igual que el wallet)
// para los logs, intentar hacer una interfaz parecido a wallet
// terminar los demas endpoints
// agregar columnas a mi tabla wallet (balance)
