package api_brasilapi

import (
	mappers "GoApi/internal/infrastructure/mappers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetCnpj fetches CNPJ data from BrasilAPI by CNPJ number.
// Returns the CNPJ data or an error on HTTP or decode failure.
func GetCnpj(cnpj string) (mappers.Cnpj, error) {
	url := fmt.Sprintf(BRASILAPI_URL, CNPJ_ENDPOINT, API_VERSION_V1, cnpj)
	response, err := http.Get(url)
	if err != nil {
		return mappers.Cnpj{}, err
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		return mappers.Cnpj{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	var cnpjData mappers.Cnpj
	err = json.NewDecoder(response.Body).Decode(&cnpjData)
	if err != nil {
		return mappers.Cnpj{}, err
	}

	return cnpjData, nil
}
