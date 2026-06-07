package webserver

import (
	"math"
	"sync"
	"time"
)

const heartbeatInterval = 60 * time.Second

type deviceRecord struct {
	heartbeats  []time.Time
	uploadTimes []time.Duration
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
	rec := s.getOrCreate(deviceID)
	rec.heartbeats = append(rec.heartbeats, t)
	s.mu.Unlock()
}

func (s *Server) RecordUploadTime(deviceID string, d time.Duration) {
	s.mu.Lock()
	rec := s.getOrCreate(deviceID)
	rec.uploadTimes = append(rec.uploadTimes, d)
	s.mu.Unlock()
}

func (s *Server) DeviceStats(deviceID string) (avgUpload string, uptime float64, hasData bool) {
	s.mu.RLock()
	rec := s.records[deviceID]
	s.mu.RUnlock()

	if rec == nil || (len(rec.heartbeats) == 0 && len(rec.uploadTimes) == 0) {
		return "", 0, false
	}
	return calcAvgUploadTime(rec.uploadTimes), calcUptime(rec.heartbeats), true
}

func (s *Server) getOrCreate(deviceID string) *deviceRecord {
	rec, ok := s.records[deviceID]
	if !ok {
		rec = &deviceRecord{}
		s.records[deviceID] = rec
	}
	return rec
}

func calcAvgUploadTime(uploads []time.Duration) string {
	if len(uploads) == 0 {
		return "0s"
	}
	var total time.Duration
	for _, d := range uploads {
		total += d
	}
	return (total / time.Duration(len(uploads))).String()
}

// calcUptime returns uptime as a percentage based on received vs expected
// heartbeats, assuming a fixed heartbeatInterval between beats.
func calcUptime(beats []time.Time) float64 {
	if len(beats) == 0 {
		return 0
	}
	if len(beats) == 1 {
		return 100
	}

	first, last := beats[0], beats[0]
	for _, t := range beats[1:] {
		if t.Before(first) {
			first = t
		}
		if t.After(last) {
			last = t
		}
	}

	elapsed := last.Sub(first)
	expected := math.Round(elapsed.Seconds()/heartbeatInterval.Seconds()) + 1
	actual := float64(len(beats))
	return math.Min(100, (actual/expected)*100)
}
