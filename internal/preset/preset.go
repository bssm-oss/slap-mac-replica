package preset

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	BundledDirectory       = "/Library/Application Support/slap-mac-replica/presets"
	HomebrewDirectory      = "/opt/homebrew/share/slap-mac-replica/presets"
	IntelHomebrewDirectory = "/usr/local/share/slap-mac-replica/presets"
	LegacyDesktopDirectory = "/Users/heodongun/Desktop/효과음"
	RandomName             = "random"
)

var supportedExtensions = map[string]bool{
	".aif":  true,
	".aiff": true,
	".m4a":  true,
	".mp3":  true,
	".wav":  true,
}

type Preset struct {
	Name string
	Path string
}

func List(dir string) ([]Preset, error) {
	if strings.TrimSpace(dir) != "" {
		return listDirectory(dir)
	}

	for _, candidate := range defaultDirectories() {
		presets, err := listDirectory(candidate)
		if err == nil && len(presets) > 0 {
			return presets, nil
		}
	}

	return nil, fmt.Errorf("no preset directory found; install presets or pass --preset-dir")
}

func listDirectory(dir string) ([]Preset, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read preset directory %s: %w", dir, err)
	}

	presets := make([]Preset, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if !supportedExtensions[ext] {
			continue
		}

		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		presets = append(presets, Preset{
			Name: name,
			Path: filepath.Join(dir, entry.Name()),
		})
	}

	sort.Slice(presets, func(i, j int) bool {
		return presets[i].Name < presets[j].Name
	})

	return presets, nil
}

func defaultDirectories() []string {
	dirs := []string{
		BundledDirectory,
		HomebrewDirectory,
		IntelHomebrewDirectory,
		LegacyDesktopDirectory,
	}

	if exe, err := os.Executable(); err == nil {
		dirs = append(dirs, filepath.Join(filepath.Dir(exe), "presets"))
	}

	return dirs
}

func Resolve(dir string, name string) (Preset, error) {
	presets, err := List(dir)
	if err != nil {
		return Preset{}, err
	}
	if len(presets) == 0 {
		return Preset{}, fmt.Errorf("no supported audio presets found in %s", dir)
	}

	if strings.EqualFold(strings.TrimSpace(name), RandomName) {
		return presets[rand.Intn(len(presets))], nil
	}

	for _, item := range presets {
		if item.Name == name || strings.EqualFold(item.Name, name) {
			return item, nil
		}
	}

	return Preset{}, fmt.Errorf("unknown preset %q; run `slap-mac-replica presets` to list available presets", name)
}
