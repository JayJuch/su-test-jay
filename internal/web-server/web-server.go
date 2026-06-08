package webserver

import (
	"log"
	"net/http"

	webhandlers "github.com/safelyyou/fleet-monitor/internal/web-server/web-handlers"
)

func (s *Server) Run(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/devices/{device_id}/heartbeat", webhandlers.Heartbeat(s))
	mux.HandleFunc("POST /api/v1/devices/{device_id}/stats", webhandlers.PostStats(s))
	mux.HandleFunc("GET /api/v1/devices/{device_id}/stats", webhandlers.GetStats(s))

	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
