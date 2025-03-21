package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gilsondev/aprendagolang-api-pgsql/models"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error parsing task ID: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(models.TaskID(taskId))
	if err != nil {
		log.Printf("Error deleting task: %v", err)
	}

	if rows > 1 {
		log.Printf("Error deleting task. Affected rows: %d", rows)
	}

	resp := map[string]any{
		"message": "Task deleted successfully",
	}

	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
