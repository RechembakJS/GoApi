package api_brasilapi

import (
	mappers "GoApi/internal/infrastructure/mappers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetAddressByBrasilApi fetches address data from BrasilAPI by zip code.
// Returns the normalized address or an error on HTTP or decode failure.
func GetAddressByBrasilApiZipcode(zipCode string) (mappers.Address, error) {
	url := fmt.Sprintf(BRASILAPI_URL, API_VERSION_V1, ZIP_CODE_ENDPOINT, zipCode)
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
