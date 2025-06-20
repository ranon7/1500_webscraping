package mediascrap

import (
	"flag"
	"fmt"
	"testing"
)

func TestValidateArgs(t *testing.T) {
	fs := flag.NewFlagSet("Test", flag.ExitOnError)
	fs.String("arg1", "value1", "")
	fs.String("arg2", "value2", "")
	fs.Parse([]string{"--arg1", "value1"})

	cases := []struct {
		actual   error
		expected error
		name     string
	}{
		{
			validateArgs([]string{"arg1"}, fs),
			nil,
			"Should not return an error",
		},
		{
			validateArgs([]string{"arg1", "arg2"}, fs),
			fmt.Errorf("missing required -%s argument", "arg2"),
			"Should return an error",
		},
	}

	for _, c := range cases {
		if errToStr(c.actual) != errToStr(c.expected) {
			t.Errorf("%s, actual %v, expected %v", c.name, c.actual, c.expected)
		}
	}
}

func errToStr(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}
