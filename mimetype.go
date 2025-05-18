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

// ExtensionByType returns the default extension for a MIME type.
// If the MIME type is not found or invalid, it returns an empty string.
func ExtensionByType(mimeType string) string {
	if mimeType == "" {
		return ""
	}

	// Convert the MIME type to lowercase to ensure case-insensitive matching
	mimeType = strings.ToLower(mimeType)

	// Retrieve the MimeType struct from the map
	if mt, exists := mimeTypes[mimeType]; exists && len(mt.Extensions) > 0 {
		// Return the first extension as the default
		return mt.Extensions[0]
	}

	// Return an empty string if no extension is found
	return ""
}
