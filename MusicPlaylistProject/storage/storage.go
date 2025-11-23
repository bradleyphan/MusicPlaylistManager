package storage

import "musicplaylist/models"

// Storage defines the interface for persistence operations
// This demonstrates Go's interface-based design and abstraction
type Storage interface {
	SavePlaylists(playlists []*models.Playlist) error
	LoadPlaylists() ([]*models.Playlist, error)
}

