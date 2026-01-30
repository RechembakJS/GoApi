// Package mappers converts external API response structs into the normalized Address type.
package mappers

import "strings"

// AddressViaCep is the ViaCEP API response shape (Portuguese field names).
type AddressViaCep struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"logradouro"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
}

// AddressBrasilApi is the BrasilAPI CEP response shape (English field names).
type AddressBrasilApi struct {
	ZipCode      string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

// Address is the normalized address used across the application (no JSON tags for external APIs).
type Address struct {
	ZipCode      string
	Street       string
	Neighborhood string
	City         string
	State        string
}

// MapAddressViaCepToAddress converts a ViaCEP response into the normalized Address.
// The zip code is normalized to digits only.
func MapAddressViaCepToAddress(address AddressViaCep) Address {
	return Address{
		ZipCode:      normalizeZipCode(address.ZipCode),
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}
}

// MapAddressBrasilApiToAddress converts a BrasilAPI response into the normalized Address.
// The zip code is normalized to digits only.
func MapAddressBrasilApiToAddress(address AddressBrasilApi) Address {
	return Address{
		ZipCode:      normalizeZipCode(address.ZipCode),
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}
}

// normalizeZipCode strips non-digit characters from the zip code (e.g. "01001-000" becomes "01001000").
func normalizeZipCode(zipCode string) string {
	var b strings.Builder
	for _, r := range zipCode {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
