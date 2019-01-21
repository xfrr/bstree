package btree

import (
	"encoding/gob"
	"os"
)

// DIRPATH ...
const DIRPATH = "storage/bst"

// Serialize serializes the object using gob
func Serialize(object interface{}) error {
	file, err := os.Create(DIRPATH)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

// Deserialize deserializes the object using gob
func Deserialize(object interface{}) error {
	file, err := os.Open(DIRPATH)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}
