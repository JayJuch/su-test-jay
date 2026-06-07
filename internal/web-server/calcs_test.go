package webserver

import (
	"testing"
	"time"
)

var base = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

func beat(offsetSeconds int) time.Time {
	return base.Add(time.Duration(offsetSeconds) * time.Second)
}

// calcUptime

func TestCalcUptime_noBeats_returnsZero(t *testing.T) {
	if got := calcUptime(time.Time{}, time.Time{}, 0); got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestCalcUptime_oneBeat_returns100(t *testing.T) {
	if got := calcUptime(base, base, 1); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestCalcUptime_noGaps_returns100(t *testing.T) {
	// 3 beats across 120s at exact 60s intervals: expected=3, actual=3
	if got := calcUptime(beat(0), beat(120), 3); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestCalcUptime_oneMissedBeat_returnsReducedPct(t *testing.T) {
	// 2 beats across 120s: expected=3, actual=2 → ~66.67%
	got := calcUptime(beat(0), beat(120), 2)
	if got < 66.6 || got > 66.7 {
		t.Fatalf("want ~66.67, got %v", got)
	}
}

func TestCalcUptime_cappedAt100(t *testing.T) {
	// More beats than expected must not exceed 100
	// 3 beats across 60s: expected=2, actual=3 → capped at 100
	if got := calcUptime(beat(0), beat(60), 3); got > 100 {
		t.Fatalf("want ≤100, got %v", got)
	}
}
