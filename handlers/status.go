package handlers

import (
	"encoding/json"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	var res map[string]any
	json.Unmarshal([]byte(`{ "status": "ok" }`), &res)

	respondWithJSON(w, http.StatusOK, res)
}
