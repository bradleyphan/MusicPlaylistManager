package main

import (
	"fmt"
	"musicplaylist/manager"
	"musicplaylist/storage"
	"musicplaylist/web"
	"os"
)

const dataFile = "playlists.json"
const port = ":8080"

func main() {
	// Initialize storage layer
	store := storage.NewJSONStorage(dataFile)

	// Initialize playlist manager
	mgr := manager.NewPlaylistManager(store)

	// Load existing data
	fmt.Println("Loading playlists...")
	err := mgr.Load()
	if err != nil {
		fmt.Printf("Warning: Could not load playlists: %v\n", err)
		fmt.Println("Starting with empty playlist collection.")
	}

	// Initialize and start web server
	server := web.NewServer(mgr, port)
	
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║   🎵 Music Playlist Manager - Web Edition  ║")
	fmt.Println("║      CS4080 Programming Languages Project  ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println()
	
	err = server.Start()
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

