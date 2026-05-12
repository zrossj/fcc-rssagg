package tools

import (
	"errors"
	"net/http"
	"strings"
)

// this should receive a api_kley, read from db and return all data related to that API_KEY.
// shoudl be in format Authorization: ApiKey {some api key here}
func GetUserAPIKEY(header http.Header) (string, error) {

	val := header.Get("Authorization")

	if val == "" {
		return "", errors.New("No auth info found") // duvida: pq usar errors.New? e quanto a log.Error?
	}

	// val e uma string no formato de um map;
	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("Malformed header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("Malformed header")
	}

	apiKey := vals[1]

	return apiKey, nil
}
