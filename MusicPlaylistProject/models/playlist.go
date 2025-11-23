package models

import (
	"fmt"
	"math/rand"
	"time"
)

// Playlist represents a collection of songs
type Playlist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Songs       []*Song   `json:"songs"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// NewPlaylist creates a new playlist
func NewPlaylist(name, description string) *Playlist {
	now := time.Now()
	return &Playlist{
		ID:          generatePlaylistID(),
		Name:        name,
		Description: description,
		Songs:       make([]*Song, 0),
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// AddSong adds a song to the playlist
func (p *Playlist) AddSong(song *Song) {
	p.Songs = append(p.Songs, song)
	p.UpdatedAt = time.Now()
}

// RemoveSong removes a song by ID from the playlist
func (p *Playlist) RemoveSong(songID string) bool {
	for i, song := range p.Songs {
		if song.ID == songID {
			p.Songs = append(p.Songs[:i], p.Songs[i+1:]...)
			p.UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

// GetSongByID retrieves a song by its ID
func (p *Playlist) GetSongByID(id string) *Song {
	for _, song := range p.Songs {
		if song.ID == id {
			return song
		}
	}
	return nil
}

// TotalDuration calculates the total duration of all songs
func (p *Playlist) TotalDuration() time.Duration {
	var total time.Duration
	for _, song := range p.Songs {
		total += song.Duration
	}
	return total
}

// Shuffle randomizes the order of songs in the playlist
func (p *Playlist) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(p.Songs), func(i, j int) {
		p.Songs[i], p.Songs[j] = p.Songs[j], p.Songs[i]
	})
	p.UpdatedAt = time.Now()
}

// String returns a formatted string representation of the playlist
func (p *Playlist) String() string {
	return fmt.Sprintf("[%s] %s - %d songs (%s total)",
		p.ID, p.Name, len(p.Songs), formatDuration(p.TotalDuration()))
}

// generatePlaylistID generates a unique playlist ID
func generatePlaylistID() string {
	return fmt.Sprintf("P%d", time.Now().UnixNano())
}

