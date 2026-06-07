package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func PostStats(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deviceID := r.PathValue("device_id")
		if !s.IsValid(deviceID) {
			writeJSON(w, http.StatusNotFound, errMsg{"device not found"})
			return
		}

		var req struct {
			SentAt     time.Time `json:"sent_at"`
			UploadTime int64     `json:"upload_time"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusInternalServerError, errMsg{err.Error()})
			return
		}

		s.RecordUploadTime(deviceID, time.Duration(req.UploadTime))
		w.WriteHeader(http.StatusNoContent)
	}
}
