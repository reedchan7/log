package log

import (
	"os"
	"strings"
)

// createFileIfNotExists creates file specified if not exists
func createFileIfNotExists(name string) error {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		// create
		return createRecursively(name)
	} else if os.IsExist(err) {
		// already exists
		return nil
	}
	return err
}

// createRecursively creates file recursively.
func createRecursively(name string) error {
	if !strings.Contains(name, "/") {
		// just a single filename
		_, err := os.Create(name)
		return err
	}

	i := strings.LastIndex(name, "/")
	path := name[:i]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	_, err := os.Create(name)
	return err
}
