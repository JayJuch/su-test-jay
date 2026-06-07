package webserver

import (
	"math"
	"sync"
	"time"
)

const heartbeatInterval = 60 * time.Second

type deviceRecord struct {
	firstBeat     time.Time
	lastBeat      time.Time
	beatCount     int
	uploadCount   int
	avgUploadTime time.Duration
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
	rec := s.getOrCreate(deviceID)
	if rec.beatCount == 0 || t.Before(rec.firstBeat) {
		rec.firstBeat = t
	}
	if rec.beatCount == 0 || t.After(rec.lastBeat) {
		rec.lastBeat = t
	}
	rec.beatCount++
}

func (s *Server) RecordUploadTime(deviceID string, d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	rec := s.getOrCreate(deviceID)
	rec.uploadCount++
	rec.avgUploadTime += (d - rec.avgUploadTime) / time.Duration(rec.uploadCount)
}

func (s *Server) DeviceStats(deviceID string) (avgUpload string, uptime float64, hasData bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	rec := s.records[deviceID]

	if rec == nil || (rec.beatCount == 0 && rec.uploadCount == 0) {
		return "", 0, false
	}
	avg := "0s"
	if rec.uploadCount > 0 {
		avg = rec.avgUploadTime.String()
	}
	return avg, calcUptime(rec.firstBeat, rec.lastBeat, rec.beatCount), true
}

func (s *Server) getOrCreate(deviceID string) *deviceRecord {
	rec, ok := s.records[deviceID]
	if !ok {
		rec = &deviceRecord{}
		s.records[deviceID] = rec
	}
	return rec
}

// calcUptime returns uptime as a percentage based on received vs expected
// heartbeats, assuming a fixed heartbeatInterval between beats.
func calcUptime(first, last time.Time, count int) float64 {
	if count == 0 {
		return 0
	}
	if count == 1 {
		return 100
	}
	elapsed := last.Sub(first)
	expected := math.Round(elapsed.Seconds()/heartbeatInterval.Seconds()) + 1
	return math.Min(100, (float64(count)/expected)*100)
}
