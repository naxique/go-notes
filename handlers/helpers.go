package handlers

import (
	"encoding/json"
	"net/http"
	Storage "notes/db"
	"notes/middleware"
)

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

type Handlers struct {
	db  *Storage.Database
	jwt *middleware.JWT
}

func (h *Handlers) SetStorage(db *Storage.Database) {
	h.db = db
}

func (h *Handlers) SetJWT(jwt *middleware.JWT) {
	h.jwt = jwt
}
