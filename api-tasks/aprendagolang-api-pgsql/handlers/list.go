package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gilsondev/aprendagolang-api-pgsql/models"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting tasks: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
