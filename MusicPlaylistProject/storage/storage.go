package storage

import "musicplaylist/models"

type Storage interface {
	SavePlaylists(playlists []*models.Playlist) error
	LoadPlaylists() ([]*models.Playlist, error)
}
