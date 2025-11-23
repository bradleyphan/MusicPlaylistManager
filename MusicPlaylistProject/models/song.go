package models

import (
	"fmt"
	"time"
)

// Song represents a music track with metadata
type Song struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Artist   string        `json:"artist"`
	Album    string        `json:"album"`
	Duration time.Duration `json:"duration"`
	Genre    string        `json:"genre"`
	Year     int           `json:"year"`
}

// NewSong creates a new song with a generated ID
func NewSong(title, artist, album string, duration time.Duration, genre string, year int) *Song {
	return &Song{
		ID:       generateID(),
		Title:    title,
		Artist:   artist,
		Album:    album,
		Duration: duration,
		Genre:    genre,
		Year:     year,
	}
}

// String returns a formatted string representation of the song
func (s *Song) String() string {
	return fmt.Sprintf("[%s] %s - %s (%s) [%s] - %s",
		s.ID, s.Title, s.Artist, s.Album, s.Genre, formatDuration(s.Duration))
}

// formatDuration converts duration to MM:SS format
func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

// generateID generates a simple unique ID (timestamp-based)
func generateID() string {
	return fmt.Sprintf("S%d", time.Now().UnixNano())
}

