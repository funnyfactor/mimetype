package mimetype

import (
	"regexp"
	"strings"
)

// facetScores defines scores for different MIME type facets
// see https://tools.ietf.org/html/rfc6838#section-3
var facetScores = map[string]float64{
	"prs.": 100,
	"x-":   200,
	"x.":   300,
	"vnd.": 400,
	"":     900, // default
}

// sourceScores defines scores for different MIME type sources
var sourceScores = map[string]float64{
	"nginx":  10,
	"apache": 20,
	"iana":   40,
	"":       30, // default
}

// typeScores defines scores for different MIME type main types
var typeScores = map[string]float64{
	// prefer application/xml over text/xml
	// prefer application/rtf over text/rtf
	"application": 1,
	// prefer font/woff over application/font-woff
	"font": 2,
	// prefer video/mp4 over application/mp4
	"video": 3,
	// default
	"": 0,
}

// facetRegex is a regex to extract the facet from the subtype
var facetRegex = regexp.MustCompile(`(\.|x-).*`)

// calculateMimeScore calculates a score for a MIME type based on its components
// The higher the score, the more "official" the type.
func calculateMimeScore(mimeType string, source string) float64 {
	if mimeType == "application/octet-stream" {
		return 0
	}

	parts := strings.Split(mimeType, "/")
	if len(parts) != 2 {
		return 0
	}

	type_, subtype := parts[0], parts[1]

	// Extract facet from subtype
	facet := facetRegex.ReplaceAllString(subtype, "$1")

	// Calculate scores
	facetScore, ok := facetScores[facet]
	if !ok {
		facetScore = facetScores[""]
	}

	sourceScore, ok := sourceScores[source]
	if !ok {
		sourceScore = sourceScores[""]
	}

	typeScore, ok := typeScores[type_]
	if !ok {
		typeScore = typeScores[""]
	}

	// All else being equal prefer shorter types
	lengthScore := 1 - float64(len(mimeType))/100

	return facetScore + sourceScore + typeScore + lengthScore
}
