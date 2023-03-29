package pathx

import (
	"errors"
	"os"
	"path/filepath"
)

func MustExpandHome(path string) string {
	res, err := ExpandHome(path)
	if err != nil {
		panic(err)
	}
	return res
}

// ExpandHome expands the path if it is start with `~`. If it isn't prefixed with `~`, then return itself
func ExpandHome(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
}

// GetHome return the home directory for current user
// if not found home directory return error
func GetHome() (string, error) {
	return os.UserHomeDir()
}
