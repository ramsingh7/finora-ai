package filemanager

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Manager performs path-safe reads and writes under a single root directory.
type Manager struct {
	root    string
	maxSize int64
}

// New validates root, resolves it to an absolute path, and ensures the directory exists.
func New(root string, maxSize int64) (*Manager, error) {
	if root == "" {
		return nil, fmt.Errorf("root: %w", ErrInvalidPath)
	}
	abs, err := filepath.Abs(filepath.Clean(root))
	if err != nil {
		return nil, fmt.Errorf("abs root: %w", err)
	}
	if err := os.MkdirAll(abs, 0o750); err != nil {
		return nil, fmt.Errorf("mkdir root: %w", err)
	}
	if maxSize <= 0 {
		return nil, fmt.Errorf("maxSize must be positive")
	}
	return &Manager{root: abs, maxSize: maxSize}, nil
}

// Root returns the absolute filesystem root for this manager.
func (m *Manager) Root() string {
	return m.root
}

// Resolve returns an absolute path under root for rel, or ErrInvalidPath if rel escapes.
func (m *Manager) Resolve(rel string) (string, error) {
	if rel == "" || strings.Contains(rel, "..") {
		return "", ErrInvalidPath
	}
	rel = filepath.ToSlash(filepath.Clean(rel))
	if strings.HasPrefix(rel, "/") || strings.HasPrefix(rel, "../") {
		return "", ErrInvalidPath
	}
	full := filepath.Join(m.root, filepath.FromSlash(rel))
	full = filepath.Clean(full)
	out, err := filepath.Rel(m.root, full)
	if err != nil || strings.HasPrefix(out, "..") {
		return "", ErrInvalidPath
	}
	return full, nil
}

// EnsureDir creates a directory tree under root (relative to rel as a dir path).
func (m *Manager) EnsureDir(rel string) error {
	p, err := m.Resolve(rel)
	if err != nil {
		return err
	}
	return os.MkdirAll(p, 0o750)
}

// WriteFile writes data to rel, creating parent directories. Total size is capped by maxSize.
func (m *Manager) WriteFile(rel string, r io.Reader) (int64, error) {
	p, err := m.Resolve(rel)
	if err != nil {
		return 0, err
	}
	if err := os.MkdirAll(filepath.Dir(p), 0o750); err != nil {
		return 0, err
	}
	tmp := p + ".tmp"
	f, err := os.OpenFile(tmp, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o640)
	if err != nil {
		return 0, err
	}
	n, err := io.Copy(f, io.LimitReader(r, m.maxSize+1))
	_ = f.Close()
	if err != nil {
		_ = os.Remove(tmp)
		return 0, err
	}
	if n > m.maxSize {
		_ = os.Remove(tmp)
		return 0, ErrTooLarge
	}
	if err := os.Rename(tmp, p); err != nil {
		_ = os.Remove(tmp)
		return 0, err
	}
	return n, nil
}

// ReadFile reads up to maxSize bytes from rel.
func (m *Manager) ReadFile(rel string) ([]byte, error) {
	p, err := m.Resolve(rel)
	if err != nil {
		return nil, err
	}
	st, err := os.Stat(p)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	if st.IsDir() {
		return nil, ErrInvalidPath
	}
	if st.Size() > m.maxSize {
		return nil, ErrTooLarge
	}
	return os.ReadFile(p)
}

// Remove deletes a file or empty directory under root.
func (m *Manager) Remove(rel string) error {
	p, err := m.Resolve(rel)
	if err != nil {
		return err
	}
	if err := os.Remove(p); err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return err
	}
	return nil
}

// Walk walks the tree under root with filepath.WalkDir semantics relative to storage root.
func (m *Manager) Walk(fn fs.WalkDirFunc) error {
	return filepath.WalkDir(m.root, fn)
}
