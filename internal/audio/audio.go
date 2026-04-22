package audio

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var builtinSounds = []string{
	"Basso",
	"Blow",
	"Bottle",
	"Frog",
	"Funk",
	"Glass",
	"Hero",
	"Morse",
	"Ping",
	"Pop",
	"Purr",
	"Sosumi",
	"Submarine",
	"Tink",
}

// Player resolves and plays either a built-in macOS sound or a custom file.
type Player struct {
	path  string
	label string
}

// NewPlayer creates a player for the requested sound name or custom file path.
func NewPlayer(value string) (Player, error) {
	path, label, err := Resolve(value)
	if err != nil {
		return Player{}, err
	}

	return Player{path: path, label: label}, nil
}

// Resolve returns the concrete file path for a built-in sound or custom file.
func Resolve(value string) (path string, label string, err error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		trimmed = "Glass"
	}

	if stat, statErr := os.Stat(trimmed); statErr == nil && !stat.IsDir() {
		abs, absErr := filepath.Abs(trimmed)
		if absErr != nil {
			return "", "", fmt.Errorf("resolve custom sound path: %w", absErr)
		}
		return abs, filepath.Base(abs), nil
	}

	normalized := normalizeBuiltin(trimmed)
	if normalized == "" {
		return "", "", fmt.Errorf("unknown sound %q; use one of %s or pass a readable file path", value, strings.Join(builtinSounds, ", "))
	}

	return filepath.Join("/System/Library/Sounds", normalized+".aiff"), normalized, nil
}

// BuiltinNames returns the built-in sound names supported by this tool.
func BuiltinNames() []string {
	names := make([]string, len(builtinSounds))
	copy(names, builtinSounds)
	return names
}

// Description returns a short human-readable sound label.
func (p Player) Description() string {
	return p.label
}

// Path returns the resolved sound path.
func (p Player) Path() string {
	return p.path
}

// Play blocks until afplay finishes.
func (p Player) Play(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "/usr/bin/afplay", p.path)
	if output, err := cmd.CombinedOutput(); err != nil {
		if len(output) == 0 {
			return fmt.Errorf("afplay %s: %w", p.path, err)
		}
		return fmt.Errorf("afplay %s: %w: %s", p.path, err, strings.TrimSpace(string(output)))
	}

	return nil
}

func normalizeBuiltin(value string) string {
	trimmed := strings.TrimSpace(value)
	trimmed = strings.TrimSuffix(trimmed, filepath.Ext(trimmed))
	if trimmed == "" {
		return ""
	}

	for _, candidate := range builtinSounds {
		if strings.EqualFold(candidate, trimmed) {
			return candidate
		}
	}

	return ""
}
