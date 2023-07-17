package path

import (
	"path/filepath"
	"runtime"
)

func Project() string {
	_, b, _, _ := runtime.Caller(0)
	currentPath := filepath.Dir(b)

	return filepath.Join(currentPath, "../../../")
}

func Migration() string {
	return filepath.Join(Project(), "database/migrations")
}
