package mediascrap

import (
	"fmt"
	"path"
	"strings"
)

func validateHref(board string, href string, acceptedFormats FileFormats) bool {
	if !strings.HasPrefix(href, fmt.Sprintf("/%s", board)) {
		return false
	}
	ext := path.Ext(href)                  // returns ".jpeg"
	format := strings.TrimPrefix(ext, ".") // returns "jpeg"

	if _, ok := acceptedFormats.Lookup[format]; !ok {
		return false
	}

	return true
}
