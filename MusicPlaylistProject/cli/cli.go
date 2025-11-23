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

// CLI handles the command-line interface
type CLI struct {
	manager *manager.PlaylistManager
	scanner *bufio.Scanner
}

// NewCLI creates a new CLI instance
func NewCLI(mgr *manager.PlaylistManager) *CLI {
	return &CLI{
		manager: mgr,
		scanner: bufio.NewScanner(os.Stdin),
	}
}

// Run starts the interactive CLI
func (c *CLI) Run() {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   🎵 Music Playlist Manager - Go Edition   ║")
	fmt.Println("║      CS4080 Programming Languages Project  ║")
	fmt.Println("╚════════════════════════════════════════════╝")
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
			fmt.Println("❌ Invalid choice. Please try again.")
		}
		fmt.Println()
	}
}

func (c *CLI) showMainMenu() {
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📋 MAIN MENU")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("1. Create New Playlist")
	fmt.Println("2. List All Playlists")
	fmt.Println("3. View Playlist Details")
	fmt.Println("4. Add Song to Playlist")
	fmt.Println("5. Remove Song from Playlist")
	fmt.Println("6. Search Songs")
	fmt.Println("7. Shuffle Playlist")
	fmt.Println("8. Delete Playlist")
	fmt.Println("9. Show Statistics")
	fmt.Println("0. Save & Exit")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

func (c *CLI) readInput(prompt string) string {
	fmt.Print(prompt)
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}

func (c *CLI) createPlaylist() {
	fmt.Println("\n➕ CREATE NEW PLAYLIST")
	name := c.readInput("Playlist name: ")
	if name == "" {
		fmt.Println("❌ Playlist name cannot be empty.")
		return
	}

	description := c.readInput("Description: ")

	playlist := c.manager.CreatePlaylist(name, description)
	fmt.Printf("✅ Created playlist: %s\n", playlist)
}

func (c *CLI) listPlaylists() {
	playlists := c.manager.ListPlaylists()

	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists found. Create one first!")
		return
	}

	fmt.Println("\n📚 ALL PLAYLISTS:")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, playlist := range playlists {
		fmt.Printf("%d. %s\n", i+1, playlist)
	}
}

func (c *CLI) viewPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	fmt.Printf("\n🎵 PLAYLIST: %s\n", playlist.Name)
	fmt.Printf("Description: %s\n", playlist.Description)
	fmt.Printf("Created: %s\n", playlist.CreatedAt.Format("2006-01-02 15:04"))
	fmt.Printf("Total Songs: %d\n", len(playlist.Songs))
	fmt.Printf("Total Duration: %v\n", formatDuration(playlist.TotalDuration()))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	if len(playlist.Songs) == 0 {
		fmt.Println("(Empty playlist)")
		return
	}

	for i, song := range playlist.Songs {
		fmt.Printf("%d. %s\n", i+1, song)
	}
}

func (c *CLI) addSongToPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists available. Create one first!")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	fmt.Println("\n🎵 ADD NEW SONG")
	title := c.readInput("Title: ")
	if title == "" {
		fmt.Println("❌ Title cannot be empty.")
		return
	}

	artist := c.readInput("Artist: ")
	album := c.readInput("Album: ")
	genre := c.readInput("Genre: ")

	durationStr := c.readInput("Duration (MM:SS): ")
	duration, err := parseDuration(durationStr)
	if err != nil {
		fmt.Printf("❌ Invalid duration format: %v\n", err)
		return
	}

	yearStr := c.readInput("Year: ")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		year = time.Now().Year()
	}

	song := models.NewSong(title, artist, album, duration, genre, year)
	playlist.AddSong(song)

	fmt.Printf("✅ Added song: %s\n", song)
}

func (c *CLI) removeSongFromPlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	if len(playlist.Songs) == 0 {
		fmt.Println("❌ Playlist is empty.")
		return
	}

	fmt.Println("\n🎵 SONGS IN PLAYLIST:")
	for i, song := range playlist.Songs {
		fmt.Printf("%d. %s\n", i+1, song)
	}

	songID := c.readInput("\nEnter song ID to remove: ")
	if playlist.RemoveSong(songID) {
		fmt.Println("✅ Song removed successfully.")
	} else {
		fmt.Println("❌ Song not found.")
	}
}

func (c *CLI) searchSongs() {
	query := c.readInput("\n🔍 Enter search query: ")
	if query == "" {
		return
	}

	results := c.manager.SearchSongs(query)

	if len(results) == 0 {
		fmt.Println("❌ No songs found matching your query.")
		return
	}

	fmt.Printf("\n✅ Found %d result(s):\n", len(results))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	for i, result := range results {
		fmt.Printf("%d. %s\n", i+1, result)
	}
}

func (c *CLI) shufflePlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID to shuffle: ")

	playlist, err := c.manager.GetPlaylist(playlistID)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	if len(playlist.Songs) == 0 {
		fmt.Println("❌ Cannot shuffle empty playlist.")
		return
	}

	playlist.Shuffle()
	fmt.Println("✅ Playlist shuffled successfully!")
}

func (c *CLI) deletePlaylist() {
	playlists := c.manager.ListPlaylists()
	if len(playlists) == 0 {
		fmt.Println("\n📭 No playlists available.")
		return
	}

	c.listPlaylists()
	playlistID := c.readInput("\nEnter playlist ID to delete: ")

	confirm := c.readInput("⚠️  Are you sure? (yes/no): ")
	if strings.ToLower(confirm) != "yes" {
		fmt.Println("❌ Deletion cancelled.")
		return
	}

	err := c.manager.DeletePlaylist(playlistID)
	if err != nil {
		fmt.Printf("❌ %v\n", err)
		return
	}

	fmt.Println("✅ Playlist deleted successfully.")
}

func (c *CLI) showStatistics() {
	stats := c.manager.GetStatistics()

	fmt.Println("\n📊 STATISTICS")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Total Playlists: %d\n", stats.TotalPlaylists)
	fmt.Printf("Total Songs: %d\n", stats.TotalSongs)
	fmt.Printf("Total Duration: %v\n", stats.TotalDuration)
	fmt.Println("\n🎸 Top Genres:")
	for genre, count := range stats.GenreCounts {
		fmt.Printf("  - %s: %d songs\n", genre, count)
	}
	fmt.Println("\n🎤 Top Artists:")
	for artist, count := range stats.ArtistCounts {
		fmt.Printf("  - %s: %d songs\n", artist, count)
	}
}

func (c *CLI) exit() {
	fmt.Println("\n💾 Saving data...")
	err := c.manager.Save()
	if err != nil {
		fmt.Printf("❌ Error saving data: %v\n", err)
	} else {
		fmt.Println("✅ Data saved successfully!")
	}
	fmt.Println("👋 Goodbye!")
}

// parseDuration parses a duration string in MM:SS format
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

// formatDuration converts duration to a readable format
func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
	}
	return fmt.Sprintf("%dm %ds", minutes, seconds)
}

