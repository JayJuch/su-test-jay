package statshandlers

import (
	"testing"
	"time"
)

var base = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

func beat(offsetSeconds int) time.Time {
	return base.Add(time.Duration(offsetSeconds) * time.Second)
}

func recordBeats(times ...time.Time) Heartbeat {
	var h Heartbeat
	for _, t := range times {
		h.Record(t)
	}
	return h
}

// Heartbeat.Uptime

func TestHeartbeat_noBeats_zeroUptime(t *testing.T) {
	var h Heartbeat
	if got := h.Uptime(); got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestHeartbeat_oneBeat_100PctUptime(t *testing.T) {
	h := recordBeats(base)
	if got := h.Uptime(); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestHeartbeat_noGaps_100PctUptime(t *testing.T) {
	// 3 beats at exact 60s intervals: expected=3, actual=3
	h := recordBeats(beat(0), beat(60), beat(120))
	if got := h.Uptime(); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestHeartbeat_oneMissedBeat_reducedUptime(t *testing.T) {
	// 2 beats across 120s: expected=3, actual=2 → ~66.67%
	h := recordBeats(beat(0), beat(120))
	got := h.Uptime()
	if got < 66.6 || got > 66.7 {
		t.Fatalf("want ~66.67, got %v", got)
	}
}

func TestHeartbeat_cappedAt100(t *testing.T) {
	// More beats than expected must not exceed 100
	h := recordBeats(beat(0), beat(30), beat(60))
	if got := h.Uptime(); got > 100 {
		t.Fatalf("want ≤100, got %v", got)
	}
}
