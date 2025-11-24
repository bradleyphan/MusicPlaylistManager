// Package storage stores playlists and songs into json on disk
package storage

import (
	"encoding/json"
	"fmt"
	"musicplaylist/models"
	"os"
)

type JSONStorage struct {
	filepath string
}

func (js *JSONStorage) SavePlaylists(playlists []*models.Playlist) error {
	data, err := json.MarshalIndent(playlists, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal playlists: %w", err)
	}

	err = os.WriteFile(js.filepath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (js *JSONStorage) LoadPlaylists() ([]*models.Playlist, error) {
	if _, err := os.Stat(js.filepath); os.IsNotExist(err) {
		return make([]*models.Playlist, 0), nil
	}

	data, err := os.ReadFile(js.filepath)
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

func NewJSONStorage(filepath string) *JSONStorage {
	return &JSONStorage{
		filepath: filepath,
	}
}
