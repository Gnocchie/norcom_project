package scanner

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/Gnocchie/norcom_project/internal/messaging"
	"github.com/Gnocchie/norcom_project/internal/model"
)

func Walk(root string, publisher messaging.Publisher) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("Skipping Walk: filepath.Walk() returned %v\n", err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		event, err := model.FromPath(path)
		if err != nil {
			log.Printf("Error reading file %v: %v\n", path, err)
			return nil
		}

		err = publisher.Publish(event)
		if err != nil {
			log.Printf("Error publishing event for %s: %v", path, err)
		}
		return nil
	})
}
