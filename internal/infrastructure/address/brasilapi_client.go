package api

import (
	mappers "GoApi/internal/infrastructure/mappers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// BRASIL_API_URL is the BrasilAPI base URL format (resource and zip code).
const BRASIL_API_URL = "https://brasilapi.com.br/api/%s/v1/%s"

// GetAddressByBrasilApi fetches address data from BrasilAPI by zip code.
// Returns the normalized address or an error on HTTP or decode failure.
func GetAddressByBrasilApi(zipCode string) (mappers.Address, error) {
	url := fmt.Sprintf(BRASIL_API_URL, "cep", zipCode)
	response, err := http.Get(url)
	if err != nil {
		return mappers.Address{}, err
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		return mappers.Address{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	var addressBrasilApi mappers.AddressBrasilApi
	err = json.NewDecoder(response.Body).Decode(&addressBrasilApi)
	if err != nil {
		return mappers.Address{}, err
	}

	return mappers.MapAddressBrasilApiToAddress(addressBrasilApi), nil
}
