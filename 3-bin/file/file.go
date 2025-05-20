package files

import (
	"path/filepath"
	"strings"
	"os"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{fileName: name}
}

func (db *JsonDb) Read() ([]byte, error) {
return os. ReadFile (db. fileName)
}

func (db *JsonDb) Write (data [ ]byte) error {
return os. WriteFile (db. fileName, data, 0644)
}

func (db *JsonDb) IsJSONFile() bool {
    db.fileName = strings.TrimSpace(db.fileName)
    extension := filepath.Ext(db.fileName)
    return strings.ToLower(extension) == ".json"
}