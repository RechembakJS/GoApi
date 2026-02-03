// Package server starts the HTTP API and registers routes for address lookup by ZIPCODE.
package server

import (
	"GoApi/internal/application/address"
	cnpjApplication "GoApi/internal/application/cnpj"
	"encoding/json"
	"log"
	"net/http"
)

// RunServer creates the HTTP mux, registers the ZIPCODE route, and starts listening on port 8080.
// It blocks until the server exits.
func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleRoot)
	mux.HandleFunc("GET /zipcode", handleZipcodeRequired)
	mux.HandleFunc("GET /zipcode/", handleZipcodeRequired)
	mux.HandleFunc("GET /zipcode/{zipcode}", handleGetAddress)
	mux.HandleFunc("GET /cnpj", handleCnpjRequired)
	mux.HandleFunc("GET /cnpj/{cnpj...}", handleGetCnpj)
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// handleRoot handles GET /. Returns 400 with usage message (path must be /zipcode/{zipcode}).
func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Use GET /zipcode/{zipcode} or GET /cnpj/{cnpj}", http.StatusBadRequest)
}

// handleZipcodeRequired handles GET /zipcode and GET /zipcode/ (missing ZIPCODE). Returns 400.
func handleZipcodeRequired(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "ZIPCODE required", http.StatusBadRequest)
}

// handleGetAddress handles GET /zipcode/{zipcode}. It looks up the address by ZIPCODE and returns JSON.
// Responds with 400 if ZIPCODE is missing, 404 if not found, and 200 with the address on success.
func handleGetAddress(w http.ResponseWriter, r *http.Request) {
	zipcode := r.PathValue("zipcode")
	if zipcode == "" {
		http.Error(w, "ZIPCODE required", http.StatusBadRequest)
		return
	}
	addr, err := address.GetAddress(zipcode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addr)
}

// handleCnpjRequired handles GET /cnpj and GET /cnpj/ (missing CNPJ). Returns 400.
func handleCnpjRequired(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "CNPJ required", http.StatusBadRequest)
}

// handleGetCnpj handles GET /cnpj/{cnpj}. It looks up the CNPJ and returns JSON.
// Responds with 400 if CNPJ is missing, 404 if not found, and 200 with the CNPJ on success.
func handleGetCnpj(w http.ResponseWriter, r *http.Request) {
	cnpj := r.PathValue("cnpj")
	if cnpj == "" {
		http.Error(w, "CNPJ required", http.StatusBadRequest)
		return
	}
	cnpjData, err := cnpjApplication.GetCnpj(cnpj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cnpjData)
}
