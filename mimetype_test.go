package mimetype

import (
	"fmt"
	"testing"
)

func TestExtensionByType(t *testing.T) {
	ext := ExtensionByType("VIDEO/MP4")
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

func TestTypeByExtension(t *testing.T) {
	mimetype := TypeByExtension(".M4A")
	if mimetype != "audio/mp4" {
		t.Errorf("expected audio/mp4, got %s", mimetype)
	}
}

func TestCalculateMimeScore(t *testing.T) {
	score := calculateMimeScore("application/vnd.ms-excel", "nginx")
	fmt.Println(score)
}
