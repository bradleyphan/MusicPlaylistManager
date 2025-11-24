// Json models for a song
package models

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/dhowden/tag"
)

type Song struct {
	ID       string        `json:"id"`
	FilePath string        `json:"filePath"`
	Title    string        `json:"title"`
	Artist   string        `json:"artist"`
	Album    string        `json:"album"`
	Duration time.Duration `json:"duration"`
	Genre    string        `json:"genre"`
	Year     int           `json:"year"`
}

func NewSongFromPath(path string, duration time.Duration) (*Song, error) {
	file, err := os.Open(path)

	if err != nil {
		return &Song{}, err
	}
	defer file.Close()

	metadata, err := tag.ReadFrom(file)

	if err != nil {
		return &Song{}, err
	}

	title := metadata.Title()

	if title == "" {
		title = filepath.Base(path)
	}

	return &Song{
		ID:       generateID(),
		Title:    title,
		Artist:   metadata.Artist(),
		Album:    metadata.Album(),
		FilePath: path,
		Genre:    metadata.Genre(),
		Year:     metadata.Year(),
		Duration: duration,
	}, nil

}

func (s *Song) ToString() string {
	return fmt.Sprintf("[%s] %s - %s (%s) [%s] - %s", s.ID, s.Title, s.Artist, s.Album, s.Genre, formatDuration(s.Duration))
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes())
	seconds := int(d.Seconds()) % 60
	return fmt.Sprintf("%d:%02d", minutes, seconds)
}

func generateID() string {
	return fmt.Sprintf("S%d", time.Now().UnixNano())
}
