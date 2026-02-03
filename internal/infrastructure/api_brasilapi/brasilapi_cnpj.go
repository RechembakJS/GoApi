package api_brasilapi

import (
	mappers "GoApi/internal/infrastructure/mappers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetCnpjV1 fetches CNPJ data from BrasilAPI by CNPJ number.
// Returns the CNPJ data or an error on HTTP or decode failure.
func GetCnpjV1(cnpj string) (mappers.CnpjV1, error) {
	url := fmt.Sprintf(BRASILAPI_URL, CNPJ_ENDPOINT, API_VERSION_V1, cnpj)
	response, err := http.Get(url)
	if err != nil {
		return mappers.CnpjV1{}, err
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		return mappers.CnpjV1{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	var cnpjV1 mappers.CnpjV1
	err = json.NewDecoder(response.Body).Decode(&cnpjV1)
	if err != nil {
		return mappers.CnpjV1{}, err
	}

	return cnpjV1, nil
}
