package audio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBuiltinSound(t *testing.T) {
	path, label, err := Resolve("glass")
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if label != "Glass" {
		t.Fatalf("expected label Glass, got %q", label)
	}

	want := filepath.Join("/System/Library/Sounds", "Glass.aiff")
	if path != want {
		t.Fatalf("expected path %q, got %q", want, path)
	}
}

func TestResolveCustomSoundPath(t *testing.T) {
	dir := t.TempDir()
	customPath := filepath.Join(dir, "custom.wav")
	if err := os.WriteFile(customPath, []byte("wave"), 0o600); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	path, label, err := Resolve(customPath)
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if path != customPath {
		t.Fatalf("expected path %q, got %q", customPath, path)
	}
	if label != "custom.wav" {
		t.Fatalf("expected label custom.wav, got %q", label)
	}
}

func TestResolveRejectsUnknownSound(t *testing.T) {
	if _, _, err := Resolve("definitely-not-a-real-sound"); err == nil {
		t.Fatal("expected Resolve to fail for unknown sound")
	}
}
