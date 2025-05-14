package bin

import "time"

type Bin struct {
	ID 			string
	Private 	bool
	CreatedAt 	time.Time
	Name 		string
}

type BinList struct {
	Bin []Bin
}

func newBin(name string, private bool) Bin {
	return Bin {
		ID: 		time.Now().Format("20060102150405"),
		Private :	private,
		CreatedAt: 	time.Now().UTC(),
		Name 	:	name,
	
	}
}
