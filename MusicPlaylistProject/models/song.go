// Json models for a song
package models

import (
	"fmt"
	"time"
)

type Song struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Artist   string        `json:"artist"`
	Album    string        `json:"album"`
	Duration time.Duration `json:"duration"`
	Genre    string        `json:"genre"`
	Year     int           `json:"year"`
}

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

func (s *Song) ToString() string {
	return fmt.Sprintf("[%s] %s - %s (%s) [%s] - %s",
		s.ID, s.Title, s.Artist, s.Album, s.Genre, formatDuration(s.Duration))
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func generateID() string {
	return fmt.Sprintf("S%d", time.Now().UnixNano())
}
