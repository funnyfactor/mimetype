package mimetype

import (
	"testing"
)

func TestExtensionByType(t *testing.T) {
	ext := ExtensionByType("video/mp4")
	if ext != "mp4" {
		t.Errorf("expected mp4, got %s", ext)
	}

	ext = ExtensionByType("application/vnd.apple.mpegurl")
	if ext != "m3u8" {
		t.Errorf("expected m3u8, got %s", ext)
	}

	ext = ExtensionByType("application/ogg")
	if ext != "ogx" {
		t.Errorf("expected ogx, got %s", ext)
	}
}
