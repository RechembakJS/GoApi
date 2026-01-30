// Package api provides HTTP clients for external CEP APIs (ViaCEP and BrasilAPI).
package api

import (
	mappers "GoApi/internal/infrastructure/mappers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// VIA_CEP_API_URL is the ViaCEP API base URL format (one %s for the zip code).
const VIA_CEP_API_URL = "https://viacep.com.br/ws/%s/json/"

// GetAddressByViaCep fetches address data from ViaCEP by zip code.
// Returns the normalized address or an error on HTTP or decode failure.
func GetAddressByViaCep(zipCode string) (mappers.Address, error) {
	url := fmt.Sprintf(VIA_CEP_API_URL, zipCode)
	response, err := http.Get(url)
	if err != nil {
		return mappers.Address{}, err
	}

	statusCode := response.StatusCode
	if statusCode != 200 {
		return mappers.Address{}, errors.New(response.Status)
	}

	defer response.Body.Close()
	var addressViaCep mappers.AddressViaCep
	err = json.NewDecoder(response.Body).Decode(&addressViaCep)
	if err != nil {
		return mappers.Address{}, err
	}

	return mappers.MapAddressViaCepToAddress(addressViaCep), nil
}
