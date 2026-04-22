package audio

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
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

const (
	// GangnamMode enables the built-in Korean speech reactions.
	GangnamMode = "gangnam"

	shortGangnamPhrase = "오빠 강남스타일"
	longGangnamPhrase  = "예~~~"
)

var (
	voiceDiscoveryOnce sync.Once
	preferredVoice     string
)

// Player resolves and plays either a built-in macOS sound, custom file, or speech phrase.
type Player struct {
	path    string
	label   string
	command string
	args    []string
}

// NewPlayer creates a player for the requested sound name, gangnam mode, or custom file path.
func NewPlayer(value string) (Player, error) {
	path, label, err := Resolve(value)
	if err != nil {
		return Player{}, err
	}

	if IsGangnamMode(value) {
		return NewGangnamShortPlayer(), nil
	}

	return Player{
		path:    path,
		label:   label,
		command: "/usr/bin/afplay",
		args:    []string{path},
	}, nil
}

// NewSpeechPlayer creates a player that speaks text with macOS say.
func NewSpeechPlayer(label string, text string) Player {
	args := make([]string, 0, 3)
	if voice := preferredKoreanVoice(); voice != "" {
		args = append(args, "-v", voice)
	}
	args = append(args, text)

	return Player{
		path:    "say:" + text,
		label:   label,
		command: "/usr/bin/say",
		args:    args,
	}
}

// NewGangnamShortPlayer returns the default short-slap speech player.
func NewGangnamShortPlayer() Player {
	return NewSpeechPlayer("gangnam-short", shortGangnamPhrase)
}

// NewGangnamLongPlayer returns the default long/rapid-slap speech player.
func NewGangnamLongPlayer() Player {
	return NewSpeechPlayer("gangnam-long", longGangnamPhrase)
}

// IsGangnamMode reports whether the current sound setting means built-in speech mode.
func IsGangnamMode(value string) bool {
	trimmed := strings.TrimSpace(value)
	return trimmed == "" || strings.EqualFold(trimmed, GangnamMode)
}

// Resolve returns the concrete file path for a built-in sound or custom file.
func Resolve(value string) (path string, label string, err error) {
	trimmed := strings.TrimSpace(value)
	if IsGangnamMode(trimmed) {
		return "say:" + shortGangnamPhrase, "gangnam-short", nil
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
		return "", "", fmt.Errorf("unknown sound %q; use %q, one of %s or pass a readable file path", value, GangnamMode, strings.Join(builtinSounds, ", "))
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

// Play blocks until the underlying playback command finishes.
func (p Player) Play(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, p.command, p.args...)
	if output, err := cmd.CombinedOutput(); err != nil {
		if len(output) == 0 {
			return fmt.Errorf("%s %s: %w", p.command, p.path, err)
		}
		return fmt.Errorf("%s %s: %w: %s", p.command, p.path, err, strings.TrimSpace(string(output)))
	}

	return nil
}

func preferredKoreanVoice() string {
	voiceDiscoveryOnce.Do(func() {
		out, err := exec.Command("/usr/bin/say", "-v", "?").Output()
		if err != nil {
			return
		}

		for _, line := range strings.Split(string(out), "\n") {
			fields := strings.Fields(line)
			if len(fields) == 0 {
				continue
			}
			if fields[0] == "Yuna" {
				preferredVoice = "Yuna"
				return
			}
		}
	})

	return preferredVoice
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
