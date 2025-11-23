# 🎵 Music Playlist Manager - Go Edition

A comprehensive music playlist management system built in Go, showcasing key programming language concepts for CS4080.

## 📚 Project Overview

This project demonstrates modern Go programming practices and core concepts from the "Concepts of Programming Languages" course, including:

- **Interfaces & Abstraction**: Storage interface for flexible persistence
- **Concurrency**: Goroutines and channels for parallel song searching
- **Synchronization**: Mutex locks for thread-safe operations
- **Error Handling**: Idiomatic Go error handling with error wrapping
- **Type System**: Strong typing with structs, methods, and custom types
- **JSON Marshaling**: Serialization/deserialization with struct tags
- **Closures**: Anonymous functions in goroutines
- **Composition**: Struct embedding and composition patterns

## 🏗️ Project Structure

```
MusicPlaylistProject/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── models/
│   ├── song.go            # Song data structure and methods
│   └── playlist.go        # Playlist data structure and operations
├── storage/
│   ├── storage.go         # Storage interface definition
│   └── json_storage.go    # JSON file storage implementation
├── manager/
│   └── playlist_manager.go # Core business logic and concurrency
├── cli/
│   └── cli.go             # Command-line interface
└── playlists.json         # Data persistence file (created at runtime)
```

## ✨ Features

### Dual Interface 🎭
- ✅ **CLI Version**: Interactive text-based menu (great for code demos)
- ✅ **Web Version**: Modern browser-based GUI (perfect for presentations)
- ✅ Both share the same backend and data file!

### Core Functionality
- ✅ Create and manage multiple playlists
- ✅ Add songs with metadata (title, artist, album, duration, genre, year)
- ✅ Remove songs from playlists
- ✅ View detailed playlist information
- ✅ Delete playlists with confirmation
- ✅ Persistent storage using JSON

### Advanced Features
- 🔍 **Concurrent Search**: Search across all playlists using goroutines
- 🔀 **Shuffle**: Randomize song order in playlists
- 📊 **Statistics**: View aggregate data (total songs, genres, artists)
- 🔒 **Thread-Safe**: All operations protected with read-write locks
- 💾 **Auto-Save**: Data persists between sessions
- 🌐 **REST API**: Full RESTful API for web interface
- 📱 **Responsive**: Web UI works on desktop, tablet, and mobile

## 🚀 Getting Started

### Prerequisites
- Go 1.21 or higher
- A terminal/command prompt
- A web browser (for web interface)

### Two Ways to Run

#### Option 1: Command-Line Interface (CLI)
```bash
go run main.go
```
Perfect for demonstrating Go concepts and code walkthrough.

#### Option 2: Web Interface (Recommended for Demos!) 🌟
```bash
go run main_web.go
```
Then open your browser to `http://localhost:8080`

Beautiful modern web interface with:
- Real-time dashboard
- Point-and-click operations
- Responsive design
- Auto-refresh

**See [WEB_GUIDE.md](WEB_GUIDE.md) for complete web interface documentation!**

### Building Executables

**CLI Version:**
```bash
go build -o playlist-manager.exe main.go
```

**Web Version:**
```bash
go build -o playlist-web.exe main_web.go
```

## 🎮 Usage Guide

### Main Menu Options

```
1. Create New Playlist      - Start a new playlist collection
2. List All Playlists       - View all your playlists
3. View Playlist Details    - See songs in a specific playlist
4. Add Song to Playlist     - Add a new song with metadata
5. Remove Song from Playlist - Delete a song from a playlist
6. Search Songs             - Find songs across all playlists
7. Shuffle Playlist         - Randomize song order
8. Delete Playlist          - Remove an entire playlist
9. Show Statistics          - View aggregate statistics
0. Save & Exit             - Save data and quit
```

### Example Workflow

1. **Create a playlist**:
   - Choose option `1`
   - Enter name: "Workout Mix"
   - Enter description: "High energy songs"

2. **Add songs**:
   - Choose option `4`
   - Enter playlist ID (shown in list)
   - Fill in song details:
     - Title: "Eye of the Tiger"
     - Artist: "Survivor"
     - Album: "Eye of the Tiger"
     - Duration: "4:05" (MM:SS format)
     - Genre: "Rock"
     - Year: "1982"

3. **Search for songs**:
   - Choose option `6`
   - Enter search term: "rock"
   - View all matching songs across playlists

4. **View statistics**:
   - Choose option `9`
   - See total playlists, songs, genres, and artists

## 🧑‍💻 Programming Concepts Demonstrated

### 1. **Interfaces** (`storage/storage.go`)
```go
type Storage interface {
    SavePlaylists(playlists []*models.Playlist) error
    LoadPlaylists() ([]*models.Playlist, error)
}
```
Demonstrates polymorphism and abstraction - easy to swap JSON storage for database storage.

### 2. **Concurrency** (`manager/playlist_manager.go`)
```go
func (pm *PlaylistManager) SearchSongs(query string) []*SearchResult {
    resultChan := make(chan *SearchResult, 100)
    var wg sync.WaitGroup
    
    for _, playlist := range pm.playlists {
        wg.Add(1)
        go func(p *models.Playlist) {
            defer wg.Done()
            // Concurrent search logic
        }(playlist)
    }
    // ...
}
```
Uses goroutines, channels, and wait groups for parallel processing.

### 3. **Synchronization** (`manager/playlist_manager.go`)
```go
type PlaylistManager struct {
    mu sync.RWMutex
    // ...
}

func (pm *PlaylistManager) GetPlaylist(id string) (*models.Playlist, error) {
    pm.mu.RLock()
    defer pm.mu.RUnlock()
    // Thread-safe read operation
}
```
Implements read-write locks for concurrent access safety.

### 4. **Error Handling**
```go
if err != nil {
    return fmt.Errorf("failed to marshal playlists: %w", err)
}
```
Uses error wrapping and propagation throughout the codebase.

### 5. **Struct Tags** (`models/song.go`)
```go
type Song struct {
    ID       string        `json:"id"`
    Title    string        `json:"title"`
    Duration time.Duration `json:"duration"`
}
```
Demonstrates metadata for JSON serialization.

### 6. **Methods on Types**
```go
func (p *Playlist) TotalDuration() time.Duration {
    var total time.Duration
    for _, song := range p.Songs {
        total += song.Duration
    }
    return total
}
```
Shows Go's approach to object-oriented programming.

## 📝 Data Persistence

Data is automatically saved to `playlists.json` when you exit the application. The file structure:

```json
[
  {
    "id": "P1234567890",
    "name": "My Playlist",
    "description": "A great collection",
    "songs": [
      {
        "id": "S9876543210",
        "title": "Example Song",
        "artist": "Example Artist",
        "album": "Example Album",
        "duration": 240000000000,
        "genre": "Rock",
        "year": 2023
      }
    ],
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:45:00Z"
  }
]
```

## 🧪 Testing the Application

### Test Scenario 1: Basic Operations
1. Create 2-3 playlists
2. Add 3-5 songs to each
3. View playlist details
4. Shuffle one playlist
5. Check that order changed

### Test Scenario 2: Search Functionality
1. Create playlists with songs from different genres
2. Search by artist name
3. Search by genre
4. Verify concurrent search works correctly

### Test Scenario 3: Persistence
1. Add several playlists and songs
2. Exit the application (option 0)
3. Restart the application
4. Verify all data is preserved

## 🎓 Key Learning Points for CS4080

1. **Type Safety**: Go's strong static typing catches errors at compile time
2. **Interface-Based Design**: Promotes loose coupling and testability
3. **Concurrency Model**: CSP-style concurrency with goroutines and channels
4. **Memory Management**: Automatic garbage collection with pointer control
5. **Error Handling**: Explicit error returns vs exception-based systems
6. **Composition**: Favors composition over inheritance
7. **Package System**: Modular code organization

## 🔧 Potential Extensions

- Add playlist import/export (CSV, M3U)
- Implement sorting (by artist, duration, year)
- Add playlist merging functionality
- Create a REST API server
- Add database storage (PostgreSQL, SQLite)
- Implement user authentication
- Add song rating system
- Support for playlist tags
- Undo/redo functionality
- Export statistics to graphs

## 📖 References

- [Go Official Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)

## 👨‍💻 Development Notes

**Language**: Go 1.21+  
**Course**: CS4080 - Concepts of Programming Languages  
**Paradigm**: Imperative, Concurrent, Statically-Typed  
**Architecture**: Layered (Models, Storage, Manager, CLI)

---

**Happy Playlist Managing! 🎵**

