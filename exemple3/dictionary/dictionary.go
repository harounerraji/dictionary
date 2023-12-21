// dictionary.go
package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Dictionary map[string]string 

func (d Dictionary) AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var entry struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&entry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	d[entry.Word] = entry.Definition
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Entry added successfully")
}

// GetHandler gère la récupération d'une définition par mot via une requête GET.
func (d Dictionary) GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	definition, exists := d[word]
	if !exists {
		http.Error(w, "Word not found", http.StatusNotFound)
		return
	}

	response := struct {
		Word       string `json:"word"`
		Definition string `json:"definition"`
	}{word, definition}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RemoveHandler gère la suppression d'une entrée par mot via une requête DELETE.
func (d Dictionary) RemoveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	_, exists := d[word]
	if !exists {
		http.Error(w, "Word not found", http.StatusNotFound)
		return
	}

	delete(d, word)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Entry removed successfully")
}

// ListHandler gère l'affichage de tous les éléments du dictionnaire via une requête GET.
func (d Dictionary) ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}
