package model

import (
	"os"
	"time"
)

type FileEvent struct {
	Path     string
	Size     int64
	ModTime  time.Time
	Hostname string
}

func FromPath(path string) (FileEvent, error) {
	info, err := os.Stat(path)
	if err != nil {
		return FileEvent{}, err
	}

	host, _ := os.Hostname()

	return FileEvent{
		Path:     path,
		Size:     info.Size(),
		ModTime:  info.ModTime(),
		Hostname: host,
	}, nil
}
