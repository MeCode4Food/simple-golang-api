package handlers

import (
	"encoding/json"
	"net/http"
	"simple-backend/server/utils/handlerutil"
)

// PingHandler Export
func PingHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	handlerutil.RespondWithJSON(w, http.StatusOK, response)
}
