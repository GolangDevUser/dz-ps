package files

import (
	"path/filepath"
	"strings"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	} 
	return data, nil
}

func IsJSONFile(fileName string) bool {
    fileName = strings.TrimSpace(fileName)
    extension := filepath.Ext(fileName)
    return strings.ToLower(extension) == ".json"
}