package config

import (
	"errors"
	"testing"
	"time"
)

func TestParseDefaultsToRun(t *testing.T) {
	cfg, err := Parse(nil)
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if cfg.Command != "run" {
		t.Fatalf("expected command run, got %q", cfg.Command)
	}
	if cfg.Run.Threshold != defaultThreshold {
		t.Fatalf("expected threshold %.2f, got %.2f", defaultThreshold, cfg.Run.Threshold)
	}
	if cfg.Run.Cooldown != defaultCooldown {
		t.Fatalf("expected cooldown %s, got %s", defaultCooldown, cfg.Run.Cooldown)
	}
	if cfg.Run.Sound != defaultSound {
		t.Fatalf("expected sound %q, got %q", defaultSound, cfg.Run.Sound)
	}
}

func TestParseRunFlags(t *testing.T) {
	cfg, err := Parse([]string{"run", "--threshold", "0.12", "--cooldown", "2s", "--sound", "Sosumi", "--short-sound", "/tmp/oppa.wav", "--rapid-sound", "/tmp/yeah.wav"})
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if cfg.Run.Threshold != 0.12 {
		t.Fatalf("expected threshold 0.12, got %.2f", cfg.Run.Threshold)
	}
	if cfg.Run.Cooldown != 2*time.Second {
		t.Fatalf("expected cooldown 2s, got %s", cfg.Run.Cooldown)
	}
	if cfg.Run.Sound != "Sosumi" {
		t.Fatalf("expected sound Sosumi, got %q", cfg.Run.Sound)
	}
	if cfg.Run.ShortSound != "/tmp/oppa.wav" {
		t.Fatalf("expected short sound path, got %q", cfg.Run.ShortSound)
	}
	if cfg.Run.RapidSound != "/tmp/yeah.wav" {
		t.Fatalf("expected rapid sound path, got %q", cfg.Run.RapidSound)
	}
}

func TestParseDoctor(t *testing.T) {
	cfg, err := Parse([]string{"doctor"})
	if err != nil {
		t.Fatalf("Parse returned error: %v", err)
	}

	if cfg.Command != "doctor" {
		t.Fatalf("expected doctor command, got %q", cfg.Command)
	}
}

func TestParseVersion(t *testing.T) {
	if _, err := Parse([]string{"version"}); !errors.Is(err, ErrVersionRequested) {
		t.Fatalf("expected ErrVersionRequested, got %v", err)
	}
}

func TestParseRejectsInvalidCooldown(t *testing.T) {
	if _, err := Parse([]string{"run", "--cooldown", "0s"}); err == nil {
		t.Fatal("expected invalid cooldown to fail")
	}
}
