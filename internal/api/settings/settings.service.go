package settings

import (
	"os"

	"github.com/gin-contrib/sessions"
)

func popSessionValues(session sessions.Session, keys ...string) (map[string]any, error) {
	values := make(map[string]any, len(keys))
	for _, key := range keys {
		values[key] = session.Get(key)
		session.Delete(key)
	}
	if err := session.Save(); err != nil {
		return nil, err
	}
	return values, nil
}

func listDirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}
