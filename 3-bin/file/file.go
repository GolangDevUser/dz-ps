package files

import (
	"path/filepath"
	"strings"
	"os"
)

func ReadFiles(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	} 
	return data, nil
}

func IsJSONFile(fileName string) bool {
	extension := filepath.Ext(fileName)
	return strings.ToLower(extension) == ".json"
}