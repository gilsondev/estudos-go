package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gilsondev/aprendagolang-api-pgsql/models"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error parsing task ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	task, err := models.Get(models.TaskID(taskId))
	if err != nil {
		log.Printf("Error getting task: %v", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
