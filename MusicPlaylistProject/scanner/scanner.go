package scanner

import (
	"errors"
	"fmt"
	"musicplaylist/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var supportedExtensions = map[string]bool{
	".mp3":  true,
	".wav":  true,
	".flac": true,
	".ogg":  true,
	".m4a":  true,
}

func ScanMusicFolder(root string) ([]*models.Song, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, errors.New("Provided path was not a directory")
	}

	var songs []*models.Song

	songChan := make(chan *models.Song)

	var wg sync.WaitGroup

	wg.Add(1)
	go walkDir(root, &wg, songChan)

	go func() {
		wg.Wait()
		close(songChan)
	}()

	for path := range songChan {
		songs = append(songs, path)
	}

	if len(songs) == 0 {
		return nil, errors.New("No songs found")
	}

	return songs, nil
}

func walkDir(dir string, wg *sync.WaitGroup, songChan chan<- *models.Song) {
	defer wg.Done()

	entries, err := os.ReadDir(dir)
	if err != nil {

		fmt.Printf("1: %v\n", err)
		return
	}

	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		if entry.IsDir() {
			wg.Add(1)
			go walkDir(fullPath, wg, songChan)
		} else {
			ext := strings.ToLower(filepath.Ext(entry.Name()))
			if supportedExtensions[ext] {
				song, err := models.NewSongFromPath(fullPath, time.Second*time.Duration(180))
				if err == nil {
					songChan <- song
				} else {

					fmt.Printf("2: %v\n", err)
				}
			}
		}
	}
}
