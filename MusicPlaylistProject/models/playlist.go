// Package models Json models for a playlist
package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Playlist struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Songs       []*Song   `json:"songs"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

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

func (p *Playlist) AddSongs(songs []*Song) {
	p.Songs = append(p.Songs, songs...)
	p.UpdatedAt = time.Now()
}

func (p *Playlist) AddSong(song *Song) {
	p.Songs = append(p.Songs, song)
	p.UpdatedAt = time.Now()
}

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

func (p *Playlist) GetSongByID(id string) *Song {
	for _, song := range p.Songs {
		if song.ID == id {
			return song
		}
	}
	return nil
}

func (p *Playlist) TotalDuration() time.Duration {
	var total time.Duration
	for _, song := range p.Songs {
		total += song.Duration
	}
	return total
}

func (p *Playlist) Shuffle() {
	rand.Shuffle(len(p.Songs), func(i, j int) {
		p.Songs[i], p.Songs[j] = p.Songs[j], p.Songs[i]
	})
	p.UpdatedAt = time.Now()
}

func (p *Playlist) ToString() string {
	return fmt.Sprintf("[%s] %s - %d songs (%s total)",
		p.ID, p.Name, len(p.Songs), formatDuration(p.TotalDuration()))
}

func generatePlaylistID() string {
	return fmt.Sprintf("P%d", time.Now().UnixNano())
}
