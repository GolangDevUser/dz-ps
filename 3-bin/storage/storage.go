package storage

import (
	"encoding/json"
	"bins/bin"
	"os"
)

func SaveBin(b bin.Bin, file string) error {
	data,err := json.Marshal(b)
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func ReadBinList(file string) (bin.BinList, error) {
	var list bin.BinList

	data, err := os.ReadFile(file)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}