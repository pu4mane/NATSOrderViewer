package pkg

import (
	"os"
	"path/filepath"
)

func ReadJSONFilesFromDirectory(dir string) ([][]byte, error) {
	var filesContent [][]byte

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			filesContent = append(filesContent, content)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesContent, nil
}
