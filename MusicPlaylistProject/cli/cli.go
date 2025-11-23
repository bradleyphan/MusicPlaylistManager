// Package cli handles CLI input
package cli

import (
	"bufio"
	"fmt"
	"musicplaylist/manager"
	"musicplaylist/models"
	"os"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	manager *manager.PlaylistManager
	scanner *bufio.Scanner
}

func CreateCLI(mgr *manager.PlaylistManager) *CLI {
	return &CLI{
		manager: mgr,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (c *CLI) Run() {
	fmt.Println("Music Playlist Manager")
	fmt.Println()

	for {
		c.showMainMenu()
		choice := c.readInput("Enter your choice: ")

		switch choice {
		case "1":
			c.createPlaylist()
		case "2":
			c.listPlaylists()
		case "3":
			c.viewPlaylist()
		case "4":
			c.addSongToPlaylist()
		case "5":
			c.removeSongFromPlaylist()
		case "6":
			c.searchSongs()
		case "7":
			c.shufflePlaylist()
		case "8":
			c.deletePlaylist()
		case "9":
			c.showStatistics()
		case "0":
			c.exit()
			return
		default:
			fmt.Println("Invalid Option")
		}
		fmt.Println()
	}
}

func (c *CLI) showMainMenu() {
	fmt.Println("MAIN MENU")
	fmt.Println("1. Create A Playlist")
	fmt.Println("2. List Playlists")
	fmt.Println("3. View Playlist Details")
	fmt.Println("4. Add Song to Playlist")
	fmt.Println("5. Remove Song from Playlist")
	fmt.Println("6. Search Songs")
	fmt.Println("7. Shuffle Playlist")
	fmt.Println("8. Delete Playlist")
	fmt.Println("9. Show Statistics")
	fmt.Println("0. Exit")
}

func (c *CLI) readInput(prompt string) string {
	fmt.Print(prompt)
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}

func (c *CLI) createPlaylist() {
	fmt.Println("\n CREATE NEW PLAYLIST")
	name := c.readInput("Playlist name: ")
	if name == "" {
		fmt.Println("Playlist name cannot be empty.")
		return
	}

	description := c.readInput("Description: ")

	playlist := c.manager.CreatePlaylist(name, description)
	fmt.Printf("Created playlist: %s\n", playlist.ToString())
}

func (c *CLI) listPlaylists() {
	playlists := c.manager.ListPlaylists()

	if len(playlists) == 0 {
		fmt.Println("\nNo playlists found.")
		return
	}

	fmt.Println("\nALL PLAYLISTS:")
	for i, playlist := range playlists {
		fmt.Printf("%d. %s\n", i+1, playlist.ToString())
	}
}

func (c *CLI) viewPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\nNo playlists found.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Printf("\nPLAYLIST: %s\n", playlist.Name)
	fmt.Printf("Description: %s\n", playlist.Description)
	fmt.Printf("Created: %s\n", playlist.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("Total Songs: %d\n", len(playlist.Songs))
	fmt.Printf("Total Duration: %v\n", durationToString(playlist.TotalDuration()))

	if len(playlist.Songs) == 0 {
		fmt.Println("(Empty playlist)")
		return
	}

	for i, song := range playlist.Songs {
		fmt.Printf("%d. %s\n", i+1, song.ToString())
	}
}

func (c *CLI) addSongToPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\nNo playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("\nADD SONG")
	title := c.readInput("Title: ")
	if title == "" {
		fmt.Println("Title cannot be empty.")
		return
	}

	artist := c.readInput("Artist: ")
	album := c.readInput("Album: ")
	genre := c.readInput("Genre: ")

	durationStr := c.readInput("Duration (MM:SS): ")
	duration, err := parseDuration(durationStr)
	if err != nil {
		fmt.Printf("Invalid duration format: %v\n", err)
		return
	}

	yearStr := c.readInput("Year: ")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = time.Now().Year()
	}

	song := models.NewSong(title, artist, album, duration, genre, year)
	playlist.AddSong(song)

	fmt.Printf("Added song: %s\n", song.ToString())
}

func (c *CLI) removeSongFromPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\nNo playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if len(playlist.Songs) == 0 {
		fmt.Println("Playlist is empty.")
		return
	}

	fmt.Println("\nSONGS IN PLAYLIST:")
	for i, song := range playlist.Songs {
		fmt.Printf("%d. %s\n", i+1, song.ToString())
	}

	songID := c.readInput("\nEnter song ID to remove: ")
	if playlist.RemoveSong(songID) {
		fmt.Println("Song removed successfully.")
	} else {
		fmt.Println("Song not found.")
	}
}

func (c *CLI) searchSongs() {
	query := c.readInput("\nEnter search query: ")
	if query == "" {
		return
	}

	results := c.manager.SearchSongs(query)

	if len(results) == 0 {
		fmt.Println("No songs found matching your query.")
		return
	}

	fmt.Printf("\nFound %d result(s):\n", len(results))
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result)
	}
}

func (c *CLI) shufflePlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\nNo playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID to shuffle: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if len(playlist.Songs) == 0 {
		fmt.Println("Playlist is empty.")
		return
	}

	playlist.Shuffle()
	fmt.Println("Playlist shuffled successfully.")
}

func (c *CLI) deletePlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\nNo playlists to delete.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID to delete: ")

	confirm := c.readInput("Are you sure? (yes/no): ")
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("Deletion cancelled.")
		return
	}

	err := c.manager.DeletePlaylist(playlistID)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	fmt.Println("Playlist deleted successfully.")
}

func (c *CLI) showStatistics() {
	stats := c.manager.GetStatistics()

	fmt.Println("\nSTATISTICS")
	fmt.Printf("Total Playlists: %d\n", stats.TotalPlaylists)
	fmt.Printf("Total Songs: %d\n", stats.TotalSongs)
	fmt.Printf("Total Duration: %v\n", stats.TotalDuration)
	fmt.Println("\nTop Genres:")
	for genre, count := range stats.GenreCounts {
		fmt.Printf("  - %s: %d songs\n", genre, count)
	}
	fmt.Println("\nTop Artists:")
	for artist, count := range stats.ArtistCounts {
		fmt.Printf("  - %s: %d songs\n", artist, count)
	}
}

func (c *CLI) exit() {
	fmt.Println("\nSaving data...")
	err := c.manager.Save()
	if err != nil {
		fmt.Printf("Error saving data: %v\n", err)
	}
	fmt.Println("Exiting.")
}

func parseDuration(s string) (time.Duration, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format, use MM:SS")
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}

func durationToString(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	return fmt.Sprintf("%dm %ds", minutes, seconds)
}
