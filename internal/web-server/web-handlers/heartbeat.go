package webhandlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func Heartbeat(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deviceID := r.PathValue("device_id")
		if !s.IsValid(deviceID) {
			writeJSON(w, http.StatusNotFound, errMsg{"device not found"})
			return
		}

		var req struct {
			SentAt time.Time `json:"sent_at"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusInternalServerError, errMsg{err.Error()})
			return
		}

		s.RecordHeartbeat(deviceID, req.SentAt)
		w.WriteHeader(http.StatusNoContent)
	}
}
