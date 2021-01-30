package finder

import (
	"os"
	"path/filepath"
)

// ExePath represent the path which process executed
func ExePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}
