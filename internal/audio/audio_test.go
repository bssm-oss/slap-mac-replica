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

func TestResolveGangnamMode(t *testing.T) {
	path, label, err := Resolve(GangnamMode)
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}

	if label != "gangnam-short" {
		t.Fatalf("expected label gangnam-short, got %q", label)
	}
	if path != "say:"+shortGangnamPhrase {
		t.Fatalf("expected speech path, got %q", path)
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

func TestIsGangnamMode(t *testing.T) {
	if !IsGangnamMode("") {
		t.Fatal("expected empty sound to enable gangnam mode")
	}
	if !IsGangnamMode("gangnam") {
		t.Fatal("expected gangnam mode to be enabled")
	}
	if IsGangnamMode("Sosumi") {
		t.Fatal("did not expect built-in afplay sound to be treated as gangnam mode")
	}
}
