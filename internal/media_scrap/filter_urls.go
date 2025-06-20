package mediascrap

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func buildUrl(relativePath string) string {
	s, _ := url.JoinPath(caravelaBaseUrl, relativePath)
	return s
}

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

func filterUrls(board string, urls []string, acceptedFormats FileFormats) []string {
	var linksFiltered []string
	for _, link := range urls {
		if strings.HasPrefix(link, fmt.Sprintf("/%s", board)) {
			mediaUrl, _ := url.JoinPath(caravelaBaseUrl, link)

			ext := path.Ext(mediaUrl)              // returns ".jpeg"
			format := strings.TrimPrefix(ext, ".") // returns "jpeg"

			if _, ok := acceptedFormats.Lookup[format]; ok {
				linksFiltered = append(linksFiltered, mediaUrl)
			}

		}
	}

	return linksFiltered
}
