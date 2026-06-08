package webhandlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Server interface {
	IsValid(deviceID string) bool
	RecordHeartbeat(deviceID string, t time.Time)
	RecordUploadTime(deviceID string, d time.Duration)
	DeviceStats(deviceID string) (avgUpload string, uptime float64, hasData bool)
}

type errMsg struct {
	Msg string `json:"msg"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
