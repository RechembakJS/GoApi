// Package cnpj provides the application use case for fetching CNPJ data by Brazilian CNPJ number.
package cnpj

import (
	api_brasilapi "GoApi/internal/infrastructure/api_brasilapi"
	mappers "GoApi/internal/infrastructure/mappers"
	"errors"
)

// GetCnpj fetches CNPJ data by CNPJ number.
// Returns the CNPJ data or an error if the API fails.
func GetCnpj(cnpj string) (mappers.Cnpj, error) {
	cnpj = mappers.CleanCnpj(cnpj)
	if len(cnpj) != 14 {
		return mappers.Cnpj{}, errors.New("CNPJ must be 14 digits")
	}
	cnpjData, err := api_brasilapi.GetCnpj(cnpj)
	if err != nil {
		return mappers.Cnpj{}, err
	}
	return cnpjData, nil
}
