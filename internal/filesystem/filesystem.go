package filesystem

import (
	"os"
	"sort"
	"time"
)

type DirectoryEntry struct {
	Name    string    `json:"name"`
	ModTime time.Time `json:"modtime"`
}

type DirectoryEntries struct {
	Directories []DirectoryEntry
	Files       []DirectoryEntry
}

func DirectoryEntriesFromPath(path string) (*DirectoryEntries, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	directories := []DirectoryEntry{}
	files := []DirectoryEntry{}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			directories = append(directories, DirectoryEntry{Name: info.Name(), ModTime: info.ModTime()})
		} else {
			files = append(files, DirectoryEntry{Name: info.Name(), ModTime: info.ModTime()})
		}
	}

	sort.Slice(directories, func(i, j int) bool {
		return directories[i].Name < directories[j].Name
	})
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	return &DirectoryEntries{Directories: directories, Files: files}, nil
}
