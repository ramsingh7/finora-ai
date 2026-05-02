package filemanager

import (
	"path/filepath"
	"testing"
)

func TestManager_ResolveBlocksTraversal(t *testing.T) {
	root := t.TempDir()
	m, err := New(root, 1024)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.Resolve("../outside")
	if err != ErrInvalidPath {
		t.Fatalf("expected ErrInvalidPath, got %v", err)
	}
}

func TestManager_Resolve_OK(t *testing.T) {
	root := t.TempDir()
	m, err := New(root, 1024)
	if err != nil {
		t.Fatal(err)
	}
	p, err := m.Resolve("safe/sub.txt")
	if err != nil {
		t.Fatalf("unexpected: %v", err)
	}
	if filepath.Base(p) != "sub.txt" {
		t.Fatalf("unexpected path %q", p)
	}
}
