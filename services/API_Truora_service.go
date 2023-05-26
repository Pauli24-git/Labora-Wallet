package services

import (
	"Labora-Wallet/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const truoraApiKey = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiIiwiYWRkaXRpb25hbF9kYXRhIjoie30iLCJjbGllbnRfaWQiOiJUQ0k4YWJkOWE1ZGFmNzM1NGQ1YjVlZjVjYTI4MjJhMjA3OSIsImV4cCI6MzI2MTY4OTIwMiwiZ3JhbnQiOiIiLCJpYXQiOjE2ODQ4ODkyMDIsImlzcyI6Imh0dHBzOi8vY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmF3cy5jb20vdXMtZWFzdC0xX3hUSGxqU1d2RCIsImp0aSI6IjM2YTZiNGJlLTM3NTUtNGQzMC04ZTM0LTNmZDMyOGI3ZDk3NCIsImtleV9uYW1lIjoidHJ1Y29kZSIsImtleV90eXBlIjoiYmFja2VuZCIsInVzZXJuYW1lIjoidHJ1b3JhdGVhbW5ld3Byb2QtdHJ1Y29kZSJ9.PuE6cS6938PbQz_4qMLySs9dr3fywFqqGdfcF6Suw0U"

const urlBase = "https://api.checks.truora.com/v1/checks/"

type API struct {
}

func (api *API) ObtainCheckID(dni int, countryId string) (string, error) {
	method := "POST"

	datos := fmt.Sprintf("national_id=%d&country=%s&type=person&user_authorized=true", dni, countryId)
	payload := strings.NewReader(datos)

	client := &http.Client{}
	req, err := http.NewRequest(method, urlBase, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Truora-API-Key", truoraApiKey)
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

	return checkID, err
}

func (api *API) ObtainScore(checkID string) (int, error) {
	method := "GET"

	newUrl := urlBase + checkID

	client := &http.Client{}
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("Truora-API-Key", truoraApiKey)
	//req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	var backgroundCheckJson models.BackgroundCheck
	err = json.NewDecoder(res.Body).Decode(&backgroundCheckJson)
	if err != nil {
		return 0, err
	}

	score := backgroundCheckJson.Check.Score

	return score, nil
}
