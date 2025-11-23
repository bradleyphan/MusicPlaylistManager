package main

import (
	"fmt"
	"musicplaylist/cli"
	"musicplaylist/manager"
	"musicplaylist/storage"
	"os"
)

const dataFile = "playlists.json"

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

	// Initialize and run CLI
	app := cli.NewCLI(mgr)
	app.Run()

	os.Exit(0)
}

