package bin

import (
	"time"

	"github.com/google/uuid"
)

type Bin struct {
	ID 			string
	Private 	bool
	CreatedAt 	time.Time
	Name 		string
}

type BinList struct {
	Bins []Bin
}

func NewBin(name string, private bool) *Bin {
	return &Bin {
		ID: 		uuid.New().String(),
		Private :	private,
		CreatedAt: 	time.Now().UTC(),
		Name 	:	name,
	
	}
}
