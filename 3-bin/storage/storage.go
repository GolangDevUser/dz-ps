package storage

import (
	"encoding/json"
	"bins/bin"
	"os"
)

type JsonStorage struct {
	fileName string
}

func NewJsonStorage(fileName string) *JsonStorage {
	return &JsonStorage{fileName: fileName}
}

func(s *JsonStorage) Read() ([]byte, error) {
	return os.ReadFile(s.fileName)
}

func(s *JsonStorage) Write(data []byte) error {
	return os.WriteFile(s.fileName, data, 0644)
}
func (s *JsonStorage) SaveBinList(list bin.BinList) error {
	data,err := json.Marshal(list)
	if err != nil {
		return err
	}
	return s.Write(data)
}

func (s *JsonStorage) ReadBinList() (bin.BinList, error) {
	var list bin.BinList

	data, err := s.Read()
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}