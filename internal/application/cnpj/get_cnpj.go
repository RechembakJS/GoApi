// Package cnpj provides the application use case for fetching CNPJ data by Brazilian CNPJ number.
package cnpj

import (
	api_brasilapi "GoApi/internal/infrastructure/api_brasilapi"
	mappers "GoApi/internal/infrastructure/mappers"
	"errors"
)

// GetCnpj fetches CNPJ data by CNPJ number.
// Returns the CNPJ data or an error if the API fails.
func GetCnpj(cnpj string) (mappers.CnpjV1, error) {
	cnpj = mappers.CleanCnpj(cnpj)
	if len(cnpj) != 14 {
		return mappers.CnpjV1{}, errors.New("CNPJ must be 14 digits")
	}
	cnpjData, err := api_brasilapi.GetCnpjV1(cnpj)
	if err != nil {
		return mappers.CnpjV1{}, err
	}
	return cnpjData, nil
}
