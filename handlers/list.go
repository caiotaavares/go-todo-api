package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/353solutions/nlp/models"
)

func List(w http.ResponseWriter, r *http.Request) {

	todo, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
