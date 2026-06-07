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
	if got := calcUptime(nil); got != 0 {
		t.Fatalf("want 0, got %v", got)
	}
}

func TestCalcUptime_oneBeat_returns100(t *testing.T) {
	if got := calcUptime([]time.Time{base}); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestCalcUptime_noGaps_returns100(t *testing.T) {
	// 3 beats at exact 60s intervals: expected=3, actual=3
	beats := []time.Time{beat(0), beat(60), beat(120)}
	if got := calcUptime(beats); got != 100 {
		t.Fatalf("want 100, got %v", got)
	}
}

func TestCalcUptime_oneMissedBeat_returnsReducedPct(t *testing.T) {
	// 2 beats across 120s: expected=3, actual=2 → ~66.67%
	beats := []time.Time{beat(0), beat(120)}
	got := calcUptime(beats)
	if got < 66.6 || got > 66.7 {
		t.Fatalf("want ~66.67, got %v", got)
	}
}

func TestCalcUptime_outOfOrderBeats_correctResult(t *testing.T) {
	// calcUptime finds min/max, so order should not matter
	beats := []time.Time{beat(60), beat(0), beat(120)}
	if got := calcUptime(beats); got != 100 {
		t.Fatalf("want 100 for out-of-order beats, got %v", got)
	}
}

func TestCalcUptime_cappedAt100(t *testing.T) {
	// More beats than expected (e.g. duplicate sends) must not exceed 100
	beats := []time.Time{beat(0), beat(30), beat(60)}
	if got := calcUptime(beats); got > 100 {
		t.Fatalf("want ≤100, got %v", got)
	}
}

// calcAvgUploadTime

func TestCalcAvgUploadTime_noUploads_returnsZeroString(t *testing.T) {
	if got := calcAvgUploadTime(nil); got != "0s" {
		t.Fatalf("want '0s', got %q", got)
	}
}

func TestCalcAvgUploadTime_singleUpload_returnsThatDuration(t *testing.T) {
	if got := calcAvgUploadTime([]time.Duration{250 * time.Millisecond}); got != "250ms" {
		t.Fatalf("want '250ms', got %q", got)
	}
}

func TestCalcAvgUploadTime_multipleUploads_returnsAverage(t *testing.T) {
	uploads := []time.Duration{100 * time.Millisecond, 300 * time.Millisecond}
	if got := calcAvgUploadTime(uploads); got != "200ms" {
		t.Fatalf("want '200ms', got %q", got)
	}
}

func TestCalcAvgUploadTime_unevenAverage_truncatesToDuration(t *testing.T) {
	// 100ms + 200ms + 300ms = 600ms / 3 = 200ms exactly
	uploads := []time.Duration{100 * time.Millisecond, 200 * time.Millisecond, 300 * time.Millisecond}
	if got := calcAvgUploadTime(uploads); got != "200ms" {
		t.Fatalf("want '200ms', got %q", got)
	}
}
