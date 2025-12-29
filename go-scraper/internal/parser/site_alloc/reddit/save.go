package reddit

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func saveRawJSON(raw string) (string, error) {
	dir := filepath.Join("cache", "reddit")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	name := fmt.Sprintf("%d.json", time.Now().UnixNano())
	path := filepath.Join(dir, name)

	err := os.WriteFile(path, []byte(raw), 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}
