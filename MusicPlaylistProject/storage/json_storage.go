package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"musicplaylist/models"
	"os"
)

// JSONStorage implements Storage interface using JSON files
type JSONStorage struct {
	filepath string
}

// NewJSONStorage creates a new JSON storage instance
func NewJSONStorage(filepath string) *JSONStorage {
	return &JSONStorage{
		filepath: filepath,
	}
}

// SavePlaylists saves playlists to a JSON file
func (js *JSONStorage) SavePlaylists(playlists []*models.Playlist) error {
	data, err := json.MarshalIndent(playlists, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal playlists: %w", err)
	}

	err = ioutil.WriteFile(js.filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// LoadPlaylists loads playlists from a JSON file
func (js *JSONStorage) LoadPlaylists() ([]*models.Playlist, error) {
	// Check if file exists
	if _, err := os.Stat(js.filepath); os.IsNotExist(err) {
		// Return empty slice if file doesn't exist
		return make([]*models.Playlist, 0), nil
	}

	data, err := ioutil.ReadFile(js.filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var playlists []*models.Playlist
	err = json.Unmarshal(data, &playlists)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal playlists: %w", err)
	}

	return playlists, nil
}

