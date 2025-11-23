package main

import (
	"fmt"
	"musicplaylist/cli"
	"musicplaylist/manager"
	"musicplaylist/storage"
	"musicplaylist/web"
	"os"
)

const dataFile = "playlists.json"
const port = ":8080"

func main() {
	store := storage.NewJSONStorage(dataFile)

	mgr := manager.CreatePlaylistManager(store)

	fmt.Println("Loading playlists...")
	err := mgr.Load()
	if err != nil {
		fmt.Printf("Warning: Could not load playlists: %v\n", err)
		fmt.Println("Starting with empty playlist collection.")
	}

	if len(os.Args) > 1 && os.Args[1] == "-web" {
		server := web.CreateServer(mgr, port)

		fmt.Println("Starting Web Server")
		fmt.Println()

		err = server.Start()
		if err != nil {
			fmt.Printf("Error starting server: %v\n", err)
			os.Exit(1)
		}

	} else {

		app := cli.CreateCLI(mgr)
		app.Run()
	}

	os.Exit(0)
}
