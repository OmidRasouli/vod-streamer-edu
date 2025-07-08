package entity

import (
	"os"
	"path/filepath"
)

// Path is a type-safe wrapper around a filesystem path string.
// This struct provides utility methods for common path operations,
// making path handling safer and more expressive throughout the codebase.
type Path struct {
	fullPath string
}

// NewPath constructs a new Path from the given parts, joining and cleaning them.
// This ensures paths are always normalized and platform-independent.
func NewPath(parts ...string) Path {
	full := filepath.Join(parts...)
	return Path{fullPath: filepath.Clean(full)}
}

// StringPathToPath converts a string path to a Path.
// Useful for adapting legacy code or external input.
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
// Useful for navigating up the directory tree.
func (p Path) Up(n int) Path {
	path := p.fullPath
	for i := 0; i < n; i++ {
		path = filepath.Dir(path)
	}
	return Path{fullPath: path}
}

// Join appends the given parts to the current path and returns a new Path.
// This helps build paths in a safe and platform-independent way.
func (p Path) Join(parts ...string) Path {
	return Path{fullPath: filepath.Join(append([]string{p.fullPath}, parts...)...)}
}

// Base returns the last element of the path.
// For example, for "/foo/bar/baz.txt" it returns "baz.txt".
func (p Path) Base() string {
	return filepath.Base(p.fullPath)
}

// IsDir returns true if the path exists and is a directory.
func (p Path) IsDir() bool {
	info, err := os.Stat(p.fullPath)
	return err == nil && info.IsDir()
}
