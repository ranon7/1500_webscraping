package mediascrap

import "testing"

func TestValidateHref(t *testing.T) {
	board := "b"
	acceptedFormats := FileFormats{
		Lookup: map[string]bool{"mp4": true},
	}

	c1actual := validateHref(board, "/b/some_file.mp4", acceptedFormats)
	c1expected := true

	c2actual := validateHref(board, "/pc/some_file.mp4", acceptedFormats)
	c2expected := false

	c3actual := validateHref(board, "https://domain/some_file.webm", acceptedFormats)
	c3expected := false

	c4actual := validateHref(board, "/b/some_file.webm", acceptedFormats)
	c4expected := false

	cases := []struct {
		name     string
		actual   bool
		expected bool
	}{
		{
			name:     "Href has the proper board prefix and file format",
			actual:   c1actual,
			expected: c1expected,
		},
		{
			name:     "Href is one of the expected formats but it's from another board",
			actual:   c2actual,
			expected: c2expected,
		},
		{
			name:     "Href has neither the right format nor the board prefix",
			actual:   c3actual,
			expected: c3expected,
		},
		{
			name:     "Href has the right board prefix but not the right format",
			actual:   c4actual,
			expected: c4expected,
		},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Fatalf("%s: actual is \n%v\nexpected: \n%v", c.name, c.actual, c.expected)
		}
	}
}
