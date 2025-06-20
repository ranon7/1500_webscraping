package mediascrap

import (
	"fmt"
	"net/url"
	"path"
)

func buildPathFromUrl(url string, folder string) string {
	fileName := path.Base(url)
	return path.Join(folder, fileName)
}

func buildThreadUrl(board string, thread string) string {
	threadHtml := fmt.Sprintf("%s.html", thread)
	s, _ := url.JoinPath(caravelaBaseUrl, board, "res", threadHtml)
	return s
}

func buildDownloadLocation(location string, board string, thread string) string {
	return path.Join(location, board, thread)
}

func buildUrl(relativePath string) string {
	s, _ := url.JoinPath(caravelaBaseUrl, relativePath)
	return s
}
