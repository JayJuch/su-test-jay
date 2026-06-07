package webserver

import (
	"log"
	"net/http"

	"github.com/safelyyou/fleet-monitor/internal/web-server/handlers"
)

func (s *Server) Run(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/devices/{device_id}/heartbeat", handlers.Heartbeat(s))
	mux.HandleFunc("POST /api/v1/devices/{device_id}/stats", handlers.PostStats(s))
	mux.HandleFunc("GET /api/v1/devices/{device_id}/stats", handlers.GetStats(s))

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
