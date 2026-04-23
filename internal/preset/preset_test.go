package preset

import (
	"os"
	"path/filepath"
	"testing"
)

func TestList(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "op-gangnam-style.mp3"))
	writeFile(t, filepath.Join(dir, "ignore.txt"))
	writeFile(t, filepath.Join(dir, "yeah.wav"))

	items, err := List(dir)
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}

	if len(items) != 2 {
		t.Fatalf("expected 2 presets, got %d", len(items))
	}
	if items[0].Name != "op-gangnam-style" {
		t.Fatalf("expected sorted first preset, got %q", items[0].Name)
	}
	if items[1].Name != "yeah" {
		t.Fatalf("expected second preset, got %q", items[1].Name)
	}
}

func TestResolve(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "op-gangnam-style.mp3")
	writeFile(t, path)

	item, err := Resolve(dir, "op-gangnam-style")
	if err != nil {
		t.Fatalf("Resolve returned error: %v", err)
	}
	if item.Path != path {
		t.Fatalf("expected path %q, got %q", path, item.Path)
	}
}

func TestResolveUnknown(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "op-gangnam-style.mp3"))

	if _, err := Resolve(dir, "missing"); err == nil {
		t.Fatal("expected unknown preset to fail")
	}
}

func writeFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("audio"), 0o600); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}
}
