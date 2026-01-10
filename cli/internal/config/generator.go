package config

import (
	"errors"
	"fmt"
	"os"
)

func Generate(path string, template string) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("%s already exists", path)
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return os.WriteFile(
		path,
		[]byte(template),
		0644,
	)
}
