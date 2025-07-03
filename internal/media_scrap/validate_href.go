package mediascrap

import (
	"path"
	"strings"
)

func validateHref(href string, acceptedFormats FileFormats) bool {
	ext := path.Ext(href)
	format := strings.TrimPrefix(ext, ".")

	if len(acceptedFormats.Lookup) > 0 {
		if _, ok := acceptedFormats.Lookup[format]; !ok {
			return false
		}
	}

	return true
}
