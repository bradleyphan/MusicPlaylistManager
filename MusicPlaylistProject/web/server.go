package web

import (
	"encoding/json"
	"fmt"
	"musicplaylist/manager"
	"musicplaylist/models"
	"net/http"
	"time"
)

// Server represents the web server
type Server struct {
	manager *manager.PlaylistManager
	port    string
}

// NewServer creates a new web server instance
func NewServer(mgr *manager.PlaylistManager, port string) *Server {
	return &Server{
		manager: mgr,
		port:    port,
	}
}

// Start starts the web server
func (s *Server) Start() error {
	// Serve static files
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/", fs)

	// API endpoints
	http.HandleFunc("/api/playlists", s.handlePlaylists)
	http.HandleFunc("/api/playlists/create", s.handleCreatePlaylist)
	http.HandleFunc("/api/playlists/delete", s.handleDeletePlaylist)
	http.HandleFunc("/api/songs/add", s.handleAddSong)
	http.HandleFunc("/api/songs/remove", s.handleRemoveSong)
	http.HandleFunc("/api/songs/search", s.handleSearchSongs)
	http.HandleFunc("/api/playlists/shuffle", s.handleShufflePlaylist)
	http.HandleFunc("/api/statistics", s.handleStatistics)

	fmt.Printf("🌐 Web server starting at http://localhost%s\n", s.port)
	fmt.Println("📱 Open your browser and visit: http://localhost" + s.port)
	fmt.Println("Press Ctrl+C to stop the server")
	
	return http.ListenAndServe(s.port, nil)
}

// handlePlaylists returns all playlists
func (s *Server) handlePlaylists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	playlists := s.manager.ListPlaylists()
	respondJSON(w, playlists)
}

// handleCreatePlaylist creates a new playlist
func (s *Server) handleCreatePlaylist(w http.ResponseWriter, r *http.Request) {
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
	
	// Save after creating
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, playlist)
}

// handleDeletePlaylist deletes a playlist
func (s *Server) handleDeletePlaylist(w http.ResponseWriter, r *http.Request) {
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

	// Save after deleting
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"status": "success"})
}

// handleAddSong adds a song to a playlist
func (s *Server) handleAddSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PlaylistID string `json:"playlist_id"`
		Title      string `json:"title"`
		Artist     string `json:"artist"`
		Album      string `json:"album"`
		Duration   int    `json:"duration"` // in seconds
		Genre      string `json:"genre"`
		Year       int    `json:"year"`
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

	song := models.NewSong(
		req.Title,
		req.Artist,
		req.Album,
		time.Duration(req.Duration)*time.Second,
		req.Genre,
		req.Year,
	)

	playlist.AddSong(song)

	// Save after adding
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, song)
}

// handleRemoveSong removes a song from a playlist
func (s *Server) handleRemoveSong(w http.ResponseWriter, r *http.Request) {
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

// handleSearchSongs searches for songs
func (s *Server) handleSearchSongs(w http.ResponseWriter, r *http.Request) {
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

// handleShufflePlaylist shuffles a playlist
func (s *Server) handleShufflePlaylist(w http.ResponseWriter, r *http.Request) {
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

	// Save after shuffling
	if err := s.manager.Save(); err != nil {
		http.Error(w, "Failed to save: "+err.Error(), http.StatusInternalServerError)
		return
	}

	respondJSON(w, playlist)
}

// handleStatistics returns statistics
func (s *Server) handleStatistics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	stats := s.manager.GetStatistics()
	respondJSON(w, stats)
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

