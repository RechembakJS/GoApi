// Package cli provides the interactive terminal menu for address lookup by zip code.
package cli

import (
	"GoApi/internal/application/address"
	cnpjApplication "GoApi/internal/application/cnpj"
	"fmt"
	"os"
)

// UserCliMenu runs the interactive menu: get address by ZIPCODE or exit.
// It prompts for a choice, then for a ZIPCODE when option 1 is selected, and prints the result.
func UserCliMenu() {
	fmt.Println("1. Get Address")
	fmt.Println("2. Get CNPJ")
	fmt.Println("0. Exit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Print("Enter the zip code: ")
		var zipCode string
		fmt.Scanln(&zipCode)
		addr, err := address.GetAddress(zipCode)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(addr)
	case 2:
		fmt.Print("Enter the CNPJ: ")
		var cnpj string
		fmt.Scanln(&cnpj)
		cnpjData, err := cnpjApplication.GetCnpj(cnpj)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(cnpjData)
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&choice)
		UserCliMenu()
	}
}
