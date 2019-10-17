package excel

import (
	"path/filepath"
	"github.com/tealeg/xlsx"
)

func ReadXlsxFile(absolutePath string) (*xlsx.File, error) {
	newFilePath := absolutePath
	if filepath.IsAbs(absolutePath) {
		newPath, _ := filepath.Abs(absolutePath)
		newFilePath = newPath
	}

	file, err := xlsx.OpenFile(newFilePath)

	if err != nil {
		return nil, err
	}

	return file, nil
}
