package mediascrap

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
)

func buildPathFromUrl(url string, folder string) string {
	fileName := path.Base(url)
	return path.Join(folder, fileName)
}

func buildThreadUrl(board string, thread int) string {
	threadHtml := fmt.Sprintf("%d.html", thread)
	s, _ := url.JoinPath(caravelaBaseUrl, board, "res", threadHtml)
	return s
}

func buildDownloadLocation(location string, board string, thread int) string {
	return path.Join(location, board, strconv.Itoa(thread))
}
