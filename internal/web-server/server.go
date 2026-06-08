package webserver

import (
	"sync"
	"time"

	statshandlers "github.com/safelyyou/fleet-monitor/internal/web-server/stats-handlers"
)

type deviceRecord struct {
	heartbeat  statshandlers.Heartbeat
	uploadTime statshandlers.UploadTime
}

type Server struct {
	validDevices map[string]struct{}
	mu           sync.RWMutex
	records      map[string]*deviceRecord
}

func NewServer(validDevices map[string]struct{}) *Server {
	return &Server{
		validDevices: validDevices,
		records:      make(map[string]*deviceRecord),
	}
}

func (s *Server) IsValid(deviceID string) bool {
	_, ok := s.validDevices[deviceID]
	return ok
}

func (s *Server) RecordHeartbeat(deviceID string, t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.getOrCreate(deviceID).heartbeat.Record(t)
}

func (s *Server) RecordUploadTime(deviceID string, d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.getOrCreate(deviceID).uploadTime.Record(d)
}

func (s *Server) DeviceStats(deviceID string) (avgUpload string, uptime float64, hasData bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rec := s.records[deviceID]

	if rec == nil || (!rec.heartbeat.HasData() && !rec.uploadTime.HasData()) {
		return "", 0, false
	}
	return rec.uploadTime.Avg(), rec.heartbeat.Uptime(), true
}

func (s *Server) getOrCreate(deviceID string) *deviceRecord {
	rec, ok := s.records[deviceID]
	if !ok {
		rec = &deviceRecord{}
		s.records[deviceID] = rec
	}
	return rec
}
