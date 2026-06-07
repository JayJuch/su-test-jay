package handlers

import "net/http"

func GetStats(s Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deviceID := r.PathValue("device_id")
		if !s.IsValid(deviceID) {
			writeJSON(w, http.StatusNotFound, errMsg{"device not found"})
			return
		}

		avgUpload, uptime, hasData := s.DeviceStats(deviceID)
		if !hasData {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"avg_upload_time": avgUpload,
			"uptime":          uptime,
		})
	}
}
