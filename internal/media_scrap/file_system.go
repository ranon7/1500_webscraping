package mediascrap

import (
	"errors"
	"fmt"
	"os"
)

func fileExists(path string) (error, bool) {
	_, err := os.Stat(path)

	if err == nil {
		return nil, true
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, false
	} else {
		return fmt.Errorf("error checking file %s", path), false
	}
}

func ensureDir(path string) error {
	err := os.MkdirAll(path, 0755) // rwxr-xr-x
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}
