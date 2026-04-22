package app

import (
	"testing"
	"time"
)

func TestUpdateRecentSlapHistoryKeepsOnlyWindow(t *testing.T) {
	now := time.Unix(10, 0)
	history := []time.Time{
		now.Add(-5 * time.Second),
		now.Add(-2 * time.Second),
		now.Add(-1 * time.Second),
	}

	updated := updateRecentSlapHistory(history, now)
	if len(updated) != 3 {
		t.Fatalf("expected 3 timestamps after trimming window, got %d", len(updated))
	}
	if updated[0] != now.Add(-2*time.Second) {
		t.Fatalf("expected oldest retained timestamp inside window, got %v", updated[0])
	}
	if updated[2] != now {
		t.Fatalf("expected newest timestamp to be current event, got %v", updated[2])
	}
}

func TestIsRapidSlapSequence(t *testing.T) {
	if isRapidSlapSequence([]time.Time{time.Now(), time.Now()}) {
		t.Fatal("did not expect two slaps to count as rapid sequence")
	}
	if !isRapidSlapSequence([]time.Time{time.Now(), time.Now(), time.Now()}) {
		t.Fatal("expected three slaps to count as rapid sequence")
	}
}
