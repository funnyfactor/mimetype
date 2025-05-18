package mimetype

import (
	_ "embed"
	"encoding/json"
	"strings"
)

type MimeType struct {
	Extensions   []string // known extensions associated with this mime type.
	Compressible bool     // whether a file of this type can be gzipped.
	Charset      string   // the default charset associated with this type, if any.
	Source       string   // where the mime type is defined. If not set, it's probably a custom media type.
}

//go:embed db.json
var db []byte // the JSON file is a map lookup for lowercased mime types

var mimeTypes = map[string]MimeType{}

func init() {
	if err := json.Unmarshal(db, &mimeTypes); err != nil {
		panic(err)
	}
}

// ExtensionByType returns the default extension for a MIME type or Content-Type value.
// If the MIME type is not found or invalid, it returns an empty string.
func ExtensionByType(mimeType string) string {
	mimeType = strings.TrimSpace(mimeType)

	// Remove any charset parameter
	if idx := strings.Index(mimeType, ";"); idx != -1 {
		mimeType = mimeType[:idx]
	}

	// If the MIME type is empty, return an empty string
	if mimeType == "" {
		return ""
	}

	// Retrieve the MimeType struct from the map
	if mt, exists := mimeTypes[strings.ToLower(mimeType)]; exists && len(mt.Extensions) > 0 {
		// Return the first extension as the default
		return mt.Extensions[0]
	}

	// Return an empty string if no extension is found
	return ""
}
