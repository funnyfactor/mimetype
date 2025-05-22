package mimetype

import (
	_ "embed"
	"encoding/json"
	"strings"
)

type mime struct {
	Extensions   []string // known extensions associated with this mime type.
	Compressible bool     // whether a file of this type can be gzipped.
	Charset      string   // the default charset associated with this type, if any.
	Source       string   // where the mime type is defined. If not set, it's probably a custom media type.
}

//go:embed db.json
var db []byte // the JSON file is a map lookup for lowercased mime types

var (
	mimeTypes = map[string]mime{}   // mime type -> mime info
	extToType = map[string]string{} // file extension -> mime type
)

func init() {
	if err := json.Unmarshal(db, &mimeTypes); err != nil {
		panic(err)
	}
	buildExtensionToMimeTypeMap()
	db = nil // free up memory
}

// buildExtensionToMimeTypeMap creates a mapping from file extensions to MIME types
func buildExtensionToMimeTypeMap() {
	for mimeType, mime := range mimeTypes {
		if len(mime.Extensions) == 0 {
			continue
		}

		for _, ext := range mime.Extensions {
			existingType, exists := extToType[ext]
			// If the extension is not mapped, or the new MIME type has a higher score, update the mapping
			if !exists || calculateMimeScore(mimeType, mime.Source) > calculateMimeScore(existingType, mimeTypes[existingType].Source) {
				extToType[ext] = mimeType
			}
		}
	}
}

// ExtensionByType returns the default extension for a MIME type or Content-Type value.
// If the MIME type is not found or invalid, it returns an empty string.
func ExtensionByType(mimeType string) string {
	// Remove charset parameter and trim spaces
	if idx := strings.Index(mimeType, ";"); idx != -1 {
		mimeType = mimeType[:idx]
	}
	mimeType = strings.ToLower(strings.TrimSpace(mimeType))
	if mimeType == "" {
		return ""
	}

	if mime := mimeTypes[mimeType]; len(mime.Extensions) > 0 {
		return mime.Extensions[0]
	}
	return ""
}

// TypeByExtension returns the MIME type for a given file extension.
// Supports both extensions with a leading dot (e.g., .mp4) and without (e.g., mp4).
// If the extension is not found or invalid, it returns an empty string.
func TypeByExtension(ext string) string {
	ext = strings.ToLower(strings.TrimPrefix(strings.TrimSpace(ext), "."))
	if ext == "" {
		return ""
	}
	return extToType[ext]
}
