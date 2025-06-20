package mediascrap

import (
	"fmt"
	"testing"
)

func TestBuildPathFromUrl(t *testing.T) {
	url := "https://1500chan.org/b/res/12132131/1232131.mp4"
	folder := "/home/anon"

	actual := buildPathFromUrl(url, folder)
	expected := "/home/anon/1232131.mp4"

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestBuildThreadUrl(t *testing.T) {
	board := "b"
	thread := "111"

	actual := buildThreadUrl(board, thread)
	expected := fmt.Sprintf("%s/b/res/111.html", caravelaBaseUrl)
	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestBuildDownloadLocation(t *testing.T) {
	location := "/home/anon/archive"
	board := "b"
	thread := "111"

	actual := buildDownloadLocation(location, board, thread)
	expected := "/home/anon/archive/b/111"

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}

func TestBuildUrl(t *testing.T) {
	relativePath := "/b/file.mp4"

	actual := buildUrl(relativePath)
	expected := fmt.Sprintf("%s/b/file.mp4", caravelaBaseUrl)

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}
