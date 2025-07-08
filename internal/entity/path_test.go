package entity

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewPath_And_String(t *testing.T) {
	p := NewPath("foo", "bar", "baz.txt")
	expected := filepath.Join("foo", "bar", "baz.txt")
	if p.String() != expected {
		t.Errorf("expected %s, got %s", expected, p.String())
	}
}

func TestParent_And_Up(t *testing.T) {
	p := NewPath("/foo/bar/baz.txt")
	parent := p.Parent()
	if parent.String() != "/foo/bar" {
		t.Errorf("expected /foo/bar, got %s", parent.String())
	}
	up := p.Up(2)
	if up.String() != "/foo" {
		t.Errorf("expected /foo, got %s", up.String())
	}
}

func TestBase(t *testing.T) {
	p := NewPath("/foo/bar/baz.txt")
	if p.Base() != "baz.txt" {
		t.Errorf("expected baz.txt, got %s", p.Base())
	}
}

func TestIsDir(t *testing.T) {
	tmpDir := t.TempDir()
	p := NewPath(tmpDir)
	if !p.IsDir() {
		t.Errorf("expected %s to be a directory", tmpDir)
	}
}

func TestValidate_Exists(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	p := NewPath(tmpFile.Name())
	if err := p.Validate(); err != nil {
		t.Errorf("expected file to validate, got error: %v", err)
	}
	if !p.Exists() {
		t.Errorf("expected file to exist")
	}
}
