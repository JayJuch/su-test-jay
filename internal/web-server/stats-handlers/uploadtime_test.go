package statshandlers

import (
	"testing"
	"time"
)

func TestUploadTime_noRecords_returnsZeroString(t *testing.T) {
	var u UploadTime
	if got := u.Avg(); got != "0s" {
		t.Fatalf("want '0s', got %q", got)
	}
}

func TestUploadTime_singleRecord_returnsThatDuration(t *testing.T) {
	var u UploadTime
	u.Record(250 * time.Millisecond)
	if got := u.Avg(); got != "250ms" {
		t.Fatalf("want '250ms', got %q", got)
	}
}

func TestUploadTime_multipleRecords_returnsAverage(t *testing.T) {
	var u UploadTime
	u.Record(100 * time.Millisecond)
	u.Record(300 * time.Millisecond)
	if got := u.Avg(); got != "200ms" {
		t.Fatalf("want '200ms', got %q", got)
	}
}
