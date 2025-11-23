package manager

import (
	"errors"
	"fmt"
	"musicplaylist/models"
	"musicplaylist/storage"
	"strings"
	"sync"
	"time"
)

// PlaylistManager manages all playlists and provides operations
type PlaylistManager struct {
	playlists []*models.Playlist
	storage   storage.Storage
	mu        sync.RWMutex // Demonstrates Go's concurrency primitives
}

// NewPlaylistManager creates a new playlist manager
func NewPlaylistManager(store storage.Storage) *PlaylistManager {
	return &PlaylistManager{
		playlists: make([]*models.Playlist, 0),
		storage:   store,
	}
}

// Load loads playlists from storage
func (pm *PlaylistManager) Load() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	playlists, err := pm.storage.LoadPlaylists()
	if err != nil {
		return err
	}

	pm.playlists = playlists
	return nil
}

// Save saves playlists to storage
func (pm *PlaylistManager) Save() error {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.storage.SavePlaylists(pm.playlists)
}

// CreatePlaylist creates a new playlist
func (pm *PlaylistManager) CreatePlaylist(name, description string) *models.Playlist {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	playlist := models.NewPlaylist(name, description)
	pm.playlists = append(pm.playlists, playlist)
	return playlist
}

// GetPlaylist retrieves a playlist by ID
func (pm *PlaylistManager) GetPlaylist(id string) (*models.Playlist, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	for _, playlist := range pm.playlists {
		if playlist.ID == id {
			return playlist, nil
		}
	}
	return nil, errors.New("playlist not found")
}

// DeletePlaylist deletes a playlist by ID
func (pm *PlaylistManager) DeletePlaylist(id string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for i, playlist := range pm.playlists {
		if playlist.ID == id {
			pm.playlists = append(pm.playlists[:i], pm.playlists[i+1:]...)
			return nil
		}
	}
	return errors.New("playlist not found")
}

// ListPlaylists returns all playlists
func (pm *PlaylistManager) ListPlaylists() []*models.Playlist {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	// Return a copy to prevent external modification
	result := make([]*models.Playlist, len(pm.playlists))
	copy(result, pm.playlists)
	return result
}

// SearchSongs searches for songs across all playlists
// This demonstrates concurrent search using goroutines and channels
func (pm *PlaylistManager) SearchSongs(query string) []*SearchResult {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	query = strings.ToLower(query)
	results := make([]*SearchResult, 0)
	resultChan := make(chan *SearchResult, 100)
	var wg sync.WaitGroup

	// Search each playlist concurrently
	for _, playlist := range pm.playlists {
		wg.Add(1)
		go func(p *models.Playlist) {
			defer wg.Done()
			for _, song := range p.Songs {
				if matchesSong(song, query) {
					resultChan <- &SearchResult{
						Song:         song,
						PlaylistName: p.Name,
						PlaylistID:   p.ID,
					}
				}
			}
		}(playlist)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}

// SearchResult represents a search result
type SearchResult struct {
	Song         *models.Song
	PlaylistName string
	PlaylistID   string
}

// String formats the search result
func (sr *SearchResult) String() string {
	return fmt.Sprintf("%s (in playlist: %s)", sr.Song.String(), sr.PlaylistName)
}

// matchesSong checks if a song matches the search query
func matchesSong(song *models.Song, query string) bool {
	return strings.Contains(strings.ToLower(song.Title), query) ||
		strings.Contains(strings.ToLower(song.Artist), query) ||
		strings.Contains(strings.ToLower(song.Album), query) ||
		strings.Contains(strings.ToLower(song.Genre), query)
}

// GetStatistics returns statistics about all playlists
func (pm *PlaylistManager) GetStatistics() Statistics {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	stats := Statistics{
		TotalPlaylists: len(pm.playlists),
		GenreCounts:    make(map[string]int),
		ArtistCounts:   make(map[string]int),
	}

	for _, playlist := range pm.playlists {
		stats.TotalSongs += len(playlist.Songs)
		stats.TotalDuration += playlist.TotalDuration()

		for _, song := range playlist.Songs {
			stats.GenreCounts[song.Genre]++
			stats.ArtistCounts[song.Artist]++
		}
	}

	return stats
}

// Statistics contains aggregate statistics
type Statistics struct {
	TotalPlaylists int
	TotalSongs     int
	TotalDuration  time.Duration
	GenreCounts    map[string]int
	ArtistCounts   map[string]int
}

