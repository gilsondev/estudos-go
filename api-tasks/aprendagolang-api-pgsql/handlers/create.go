package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gilsondev/aprendagolang-api-pgsql/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(task)

	var resp map[string]any
	if err != nil {
		log.Printf("Error inserting task: %v", err)
		errMsg := fmt.Sprintf("Error inserting task: %v", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	resp = map[string]any{
		"message": "Task created successfully",
		"id":      id,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
