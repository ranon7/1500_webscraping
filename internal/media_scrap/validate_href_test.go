package mediascrap

import "testing"

func TestValidateHref(t *testing.T) {
	acceptedFormats := FileFormats{
		Lookup: map[string]bool{"mp4": true},
	}

	c1actual := validateHref("/b/some_file.mp4", acceptedFormats)
	c1expected := true

	c2actual := validateHref("/b/some_file.webm", acceptedFormats)
	c2expected := false

	cases := []struct {
		name     string
		actual   bool
		expected bool
	}{
		{
			name:     "Href has the proper file format",
			actual:   c1actual,
			expected: c1expected,
		},
		{
			name:     "Href has not the right format",
			actual:   c2actual,
			expected: c2expected,
		},
	}

	for _, c := range cases {
		if c.actual != c.expected {
			t.Fatalf("%s: actual is %v expected %v", c.name, c.actual, c.expected)
		}
	}
}
