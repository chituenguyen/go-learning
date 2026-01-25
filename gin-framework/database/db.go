package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type JSONDatabase struct {
	filepath string
	mu       sync.RWMutex
}

func NewJSONDatabase(filepath string) *JSONDatabase {
	return &JSONDatabase{
		filepath: filepath,
	}
}

// ReadData reads and unmarshals JSON data from file
func (db *JSONDatabase) ReadData(v interface{}) error {
	db.mu.RLock()
	defer db.mu.RUnlock()

	// Check if file exists
	if _, err := os.Stat(db.filepath); os.IsNotExist(err) {
		// Create empty file with empty array
		return db.WriteData([]interface{}{})
	}

	data, err := ioutil.ReadFile(db.filepath)
	if err != nil {
		return err
	}

	// If file is empty, initialize with empty array
	if len(data) == 0 {
		data = []byte("[]")
	}

	return json.Unmarshal(data, v)
}

// WriteData marshals and writes data to JSON file
func (db *JSONDatabase) WriteData(v interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(db.filepath, data, 0644)
}
