package db

import (
	"encoding/gob"
	"io"
)

// SaveToFile writes a database to a file
func SaveToFile(store *Database, writer io.Writer) error {
	encoder := gob.NewEncoder(writer)

	return encoder.Encode(*store)
}

// ReadFromFile reads from a file to server
func ReadFromFile(reader io.Reader) (*Database, error) {
	db := NewDb("")

	decoder := gob.NewDecoder(reader)
	if err := decoder.Decode(db); err != nil {
		return nil, err
	}

	return db, nil
}
