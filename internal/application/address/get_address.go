// Package address provides the application use case for fetching address data by Brazilian zip code (CEP).
package address

import (
	api "GoApi/internal/infrastructure/address"
	"GoApi/internal/infrastructure/mappers"
)

// GetAddress fetches an address by zip code (CEP). It tries ViaCEP first, then BrasilAPI as fallback.
// Returns the normalized address or an error if both providers fail.
func GetAddress(zipCode string) (mappers.Address, error) {
	address, err := api.GetAddressByViaCep(zipCode)
	if err != nil {
		address, err = api.GetAddressByBrasilApi(zipCode)
		if err != nil {
			return mappers.Address{}, err
		}
	}

	return address, nil
}
