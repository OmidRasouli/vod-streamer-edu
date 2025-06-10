package entity

import (
	"os"
	"path/filepath"
)

// Path is a type-safe wrapper around a filesystem path string.
type Path struct {
	fullPath string
}

// NewPath constructs a new Path from the given parts, joining and cleaning them.
func NewPath(parts ...string) Path {
	full := filepath.Join(parts...)
	return Path{fullPath: filepath.Clean(full)}
}

// StringPathToPath converts a string path to a Path.
func StringPathToPath(path string) Path {
	return NewPath(path)
}

// Validate checks if the path is non-empty and exists in the filesystem.
// Returns an error if the path is invalid or does not exist.
func (p Path) Validate() error {
	if p.fullPath == "" {
		return os.ErrInvalid
	}
	_, err := os.Stat(p.fullPath)
	return err
}

// Exists returns true if the path exists in the filesystem.
func (p Path) Exists() bool {
	_, err := os.Stat(p.fullPath)
	return err == nil
}

// String returns the underlying string representation of the path.
func (p Path) String() string {
	return p.fullPath
}

// Parent returns the parent directory as a new Path.
func (p Path) Parent() Path {
	return Path{fullPath: filepath.Dir(p.fullPath)}
}

// Up returns a new Path that is n directories above the current path.
func (p Path) Up(n int) Path {
	path := p.fullPath
	for i := 0; i < n; i++ {
		path = filepath.Dir(path)
	}
	return Path{fullPath: path}
}

// Join appends the given parts to the current path and returns a new Path.
func (p Path) Join(parts ...string) Path {
	return Path{fullPath: filepath.Join(append([]string{p.fullPath}, parts...)...)}
}

// Base returns the last element of the path.
func (p Path) Base() string {
	return filepath.Base(p.fullPath)
}

// IsDir returns true if the path exists and is a directory.
func (p Path) IsDir() bool {
	info, err := os.Stat(p.fullPath)
	return err == nil && info.IsDir()
}
