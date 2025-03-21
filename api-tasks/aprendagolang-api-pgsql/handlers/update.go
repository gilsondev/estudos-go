package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gilsondev/aprendagolang-api-pgsql/models"
)

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error parsing task ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var task models.Task

	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Update(models.TaskID(taskId), task)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error updating task. Affected rows: %d", rows)
	}

	resp := map[string]any{
		"message": "Task updated successfully",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
