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
	threadHtml := fmt.Sprintf("%s.json", thread)
	s, _ := url.JoinPath(caravelaBaseUrl, board, "res", threadHtml)
	return s
}

func buildFileUrl(board string, filename string, ext string) string {
	f := fmt.Sprintf("%s%s", filename, ext)
	s, _ := url.JoinPath(caravelaBaseUrl, board, "src", f)

	return s
}

func buildDownloadLocation(location string, board string, thread string) string {
	return path.Join(location, board, thread)
}
