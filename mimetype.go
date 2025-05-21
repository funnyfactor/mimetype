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

var mimes = map[string]mime{}                 // mime type -> mime
var extensionToMimeType = map[string]string{} // extension -> mime type

func init() {
	if err := json.Unmarshal(db, &mimes); err != nil {
		panic(err)
	}
	buildExtensionToMimeTypeMap()
	db = nil // free up memory
}

// ExtensionByType returns the default extension for a MIME type or Content-Type value.
// If the MIME type is not found or invalid, it returns an empty string.
func ExtensionByType(mimeType string) string {
	// Remove charset parameter and trim spaces
	if idx := strings.Index(mimeType, ";"); idx != -1 {
		mimeType = mimeType[:idx]
	}
	mimeType = strings.TrimSpace(strings.ToLower(mimeType))
	if mimeType == "" {
		return ""
	}

	if mime := mimes[mimeType]; len(mime.Extensions) > 0 {
		return mime.Extensions[0]
	}
	return ""
}

// TypeByExtension returns the MIME type for a given file extension.
// If the extension is not found or invalid, it returns an empty string.
func TypeByExtension(ext string) string {
	ext = strings.TrimSpace(ext)
	if ext == "" {
		return ""
	}

	// Remove leading dot if present
	ext = strings.TrimPrefix(ext, ".")
	return extensionToMimeType[ext]
}

// buildExtensionToMimeTypeMap creates a mapping from file extensions to MIME types
func buildExtensionToMimeTypeMap() {
	for mimeType, mime := range mimes {
		if len(mime.Extensions) == 0 {
			continue
		}

		for _, ext := range mime.Extensions {
			// If the extension is already mapped, check which MIME type has higher score
			if existingType, exists := extensionToMimeType[ext]; exists {
				existingScore := calculateMimeScore(existingType, mimes[existingType].Source)
				newScore := calculateMimeScore(mimeType, mime.Source)

				// If new type has higher score, update the mapping
				if newScore > existingScore {
					extensionToMimeType[ext] = mimeType
				}
			} else {
				extensionToMimeType[ext] = mimeType
			}
		}
	}
}
