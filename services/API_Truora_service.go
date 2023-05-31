package services

import (
	"Labora-Wallet/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type API struct {
	apiKey        string
	urlBaseTruora string
}

func (api *API) initConfig() error {
	var err error

	config, err := models.LoadEnvVariablesTruora()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	api.apiKey = config.TruoraApiKey
	api.urlBaseTruora = config.Urlbase

	return err
}

func (api *API) ObtainCheckID(dni int, countryId string) (string, error) {
	method := "POST"

	payload := url.Values{
		"national_id":     {fmt.Sprintf("%d", dni)},
		"country":         {countryId},
		"type":            {"person"},
		"user_authorized": {"true"},
	}.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(method, api.urlBaseTruora, strings.NewReader(payload))
	if err != nil {
		return "", err
	}

	req.Header.Add("Truora-API-Key", api.apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var backgroundCheck models.BackgroundCheck
	err = json.NewDecoder(res.Body).Decode(&backgroundCheck)
	if err != nil {
		return "", err
	}

	checkID := backgroundCheck.Check.CheckID
	if checkID == "" {
		return "", errors.New("El checkId no puede estar vacio")
	}
	return checkID, err
}

func (api *API) ObtainScore(checkID string) (int, error) {
	var finish bool
	var score int

	method := "GET"

	newUrl := api.urlBaseTruora + checkID

	client := &http.Client{}
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Truora-API-Key", api.apiKey)

	for !finish {

		res, err := client.Do(req)

		if err != nil {
			break
		}

		var backgroundCheckJson models.BackgroundCheck

		err = json.NewDecoder(res.Body).Decode(&backgroundCheckJson)

		if err != nil {
			break
		}

		if backgroundCheckJson.Check.Status == "completed" {
			score = backgroundCheckJson.Check.Score
			finish = true
		}

		if backgroundCheckJson.Check.Status == "error" {
			finish = true
		}

	}

	if err != nil {
		return 0, err
	}

	return score, nil
}
