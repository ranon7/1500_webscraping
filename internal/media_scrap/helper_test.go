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
	expected := fmt.Sprintf("%s/b/res/111.json", caravelaBaseUrl)
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

func TestBuildFileUrl(t *testing.T) {
	board := "b"
	filename := "myfile"
	ext := ".mp4"

	actual := buildFileUrl(board, filename, ext)
	expected := fmt.Sprintf("%s/%s/src/%s%s", caravelaBaseUrl, board, filename, ext)

	if actual != expected {
		t.Fatalf("actual %s expected %s", actual, expected)
	}
}
