package main

import (
	"bins/bin"
	"bins/storage"
	"encoding/json"
	"fmt"
)

func main() {
	// var db bin.Db = files.NewJsonDb("bins.json")
	var db bin.Db = storage.NewJsonStorage("storageBin.json")
	binList := bin.BinList{
		Bins: []bin.Bin{
			*bin.NewBin("Bin1", false),
			*bin.NewBin("Bin2", true),
			*bin.NewBin("Bin3", false),
			*bin.NewBin("Bin4", true),
		},
	}

	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	err = db.Write(data)
	if err != nil {
		fmt.Println("Ошибка записи в файл", err)
		return
	}

	readData, err := db.Read()
	if err != nil {
		fmt.Println("Ошибка чтения из файла: ", err)
		return
	}

	var readList bin.BinList
	err = json.Unmarshal(readData, &readList)
	if err != nil {
		fmt.Println("Ошибка")
		return
	}
	fmt.Println("Прочитанный список: ", readList)
}