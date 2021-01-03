package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"simple-backend/server/models/cat"
)

// DB Export
type DB struct {
	CatStore []cat.Cat // not really a "database" per se
}

// CreateCatDatabase Export
func CreateCatDatabase() DB {
	db := DB{CatStore: make([]cat.Cat, 0)}

	return db
}

// InitialiseSampleData Export
func (db *DB) InitialiseSampleData() {

	// path of file is relative to the location of main.go

	// Read json from sample-data.json
	// for each item inside list, load into CatStore
	data, err := ioutil.ReadFile("./server/services/db/sample-data.json")

	var cats []cat.Cat

	json.Unmarshal(data, &cats)
	if err != nil {
		panic(err)
	}
	for _, cat := range cats {
		db.CatStore = append(db.CatStore, cat)
	}
	fmt.Println(cats)
}
