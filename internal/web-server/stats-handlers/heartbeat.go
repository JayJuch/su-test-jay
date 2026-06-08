package statshandlers

import (
	"math"
	"time"
)

const HeartbeatInterval = 60 * time.Second

type Heartbeat struct {
	firstBeat time.Time
	lastBeat  time.Time
	count     int
}

func (h *Heartbeat) Record(t time.Time) {
	if h.count == 0 || t.Before(h.firstBeat) {
		h.firstBeat = t
	}
	if h.count == 0 || t.After(h.lastBeat) {
		h.lastBeat = t
	}
	h.count++
}

func (h *Heartbeat) HasData() bool {
	return h.count > 0
}

// Uptime returns uptime as a percentage based on received vs expected
// heartbeats, assuming a fixed HeartbeatInterval between beats.
func (h *Heartbeat) Uptime() float64 {
	return calcUptime(h.firstBeat, h.lastBeat, h.count)
}

func calcUptime(first, last time.Time, count int) float64 {
	if count == 0 {
		return 0
	}
	if count == 1 {
		return 100
	}
	elapsed := last.Sub(first)
	expected := math.Round(elapsed.Seconds()/HeartbeatInterval.Seconds()) + 1
	return math.Min(100, (float64(count)/expected)*100)
}
