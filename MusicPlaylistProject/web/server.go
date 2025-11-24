// Package web is the webserver implementation
package web

import (
	"encoding/json"
	"fmt"
	"musicplaylist/manager"
	"musicplaylist/models"
	"musicplaylist/scanner"
	"net/http"
	"time"
)

type WebServer struct {
	manager *manager.PlaylistManager
	port    string
}

func CreateServer(mgr *manager.PlaylistManager, port string) *WebServer {
	return &WebServer{
		manager: mgr,
		port:    port,
	}
}

func (s *WebServer) Start() error {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/playlists", s.handlePlaylists)
	http.HandleFunc("/api/playlists/create", s.handleCreatePlaylist)
	http.HandleFunc("/api/playlists/delete", s.handleDeletePlaylist)
	http.HandleFunc("/api/songs/add", s.handleAddSong)

	http.HandleFunc("/api/songs/scan", s.handleScanFolder)
	http.HandleFunc("/api/songs/remove", s.handleRemoveSong)
	http.HandleFunc("/api/songs/search", s.handleSearchSongs)
	http.HandleFunc("/api/playlists/shuffle", s.handleShufflePlaylist)
	http.HandleFunc("/api/statistics", s.handleStatistics)

	fmt.Printf("Web server starting at http://localhost%s\n", s.port)
	fmt.Println("Press Ctrl+C to stop the server")

	return http.ListenAndServe(s.port, nil)
}

func (s *WebServer) handlePlaylists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playlists := s.manager.ListPlaylists()
	respondJSON(w, playlists)
}

func (s *WebServer) handleCreatePlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist := s.manager.CreatePlaylist(req.Name, req.Description)

	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, playlist)
}

func (s *WebServer) handleDeletePlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.manager.DeletePlaylist(req.ID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "success"})
}

func (s *WebServer) handleScanFolder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		FilePath   string `json:"file_path"`
		PlaylistID string `json:"playlist_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := s.manager.GetPlaylist(req.PlaylistID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	songs, err := scanner.ScanMusicFolder(req.FilePath)

	if err != nil {
		http.Error(w, "Error Scanning Folder: "+err.Error(), http.StatusInternalServerError)

		return
	}

	playlist.AddSongs(songs)

	// Save after adding
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, songs)
}

func (s *WebServer) handleAddSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		FilePath   string `json:"file_path"`
		PlaylistID string `json:"playlist_id"`
		Duration   int    `json:"duration"` // in seconds
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := s.manager.GetPlaylist(req.PlaylistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	song, err := models.NewSongFromPath(req.FilePath,
		time.Duration(req.Duration)*time.Second,
	)
	if err != nil {
		http.Error(w, "Failed Add Song: "+err.Error(), http.StatusInternalServerError)
		return
	}

	playlist.AddSong(song)

	// Save after adding
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, song)
}

func (s *WebServer) handleRemoveSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PlaylistID string `json:"playlist_id"`
		SongID     string `json:"song_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := s.manager.GetPlaylist(req.PlaylistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !playlist.RemoveSong(req.SongID) {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	// Save after removing
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "success"})
}

func (s *WebServer) handleSearchSongs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing search query", http.StatusBadRequest)
		return
	}

	results := s.manager.SearchSongs(query)
	respondJSON(w, results)
}

func (s *WebServer) handleShufflePlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := s.manager.GetPlaylist(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	playlist.Shuffle()

	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, playlist)
}

func (s *WebServer) handleStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats := s.manager.GetStatistics()
	respondJSON(w, stats)
}

func respondJSON(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		println("Error Encoding JSON")
	}
}
