# 🏗️ Architecture Documentation

## System Architecture

The Music Playlist Manager follows a **layered architecture** pattern with clear separation of concerns:

```
┌─────────────────────────────────────┐
│         CLI Layer (cli/)            │  ← User Interface
│   - Interactive menu system         │
│   - Input/output handling           │
└────────────┬────────────────────────┘
             │
             ▼
┌─────────────────────────────────────┐
│    Business Logic (manager/)        │  ← Core Operations
│   - Playlist management             │
│   - Concurrent search               │
│   - Thread-safe operations          │
└────────────┬────────────────────────┘
             │
             ▼
┌─────────────────────────────────────┐
│      Storage Layer (storage/)       │  ← Persistence
│   - Interface abstraction           │
│   - JSON implementation             │
└────────────┬────────────────────────┘
             │
             ▼
┌─────────────────────────────────────┐
│      Data Models (models/)          │  ← Data Structures
│   - Song struct & methods           │
│   - Playlist struct & methods       │
└─────────────────────────────────────┘
```

## Component Breakdown

### 1. Data Models Layer (`models/`)

**Purpose**: Define core data structures and their behaviors

#### Song (`models/song.go`)
- **Responsibility**: Represents a music track with metadata
- **Key Features**:
  - Immutable ID generation
  - Duration formatting
  - String representation
- **Design Pattern**: Value Object

```go
type Song struct {
    ID       string
    Title    string
    Artist   string
    Album    string
    Duration time.Duration
    Genre    string
    Year     int
}
```

#### Playlist (`models/playlist.go`)
- **Responsibility**: Collection of songs with operations
- **Key Features**:
  - Song management (add/remove)
  - Total duration calculation
  - Shuffle algorithm
  - Timestamp tracking
- **Design Pattern**: Aggregate Root

```go
type Playlist struct {
    ID          string
    Name        string
    Description string
    Songs       []*Song
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### 2. Storage Layer (`storage/`)

**Purpose**: Abstract persistence mechanism

#### Storage Interface (`storage/storage.go`)
- **Responsibility**: Define persistence contract
- **Design Pattern**: Interface Segregation Principle

```go
type Storage interface {
    SavePlaylists(playlists []*models.Playlist) error
    LoadPlaylists() ([]*models.Playlist, error)
}
```

**Benefits**:
- Easy to test with mock implementations
- Can swap JSON for database without changing business logic
- Dependency Inversion Principle

#### JSON Storage (`storage/json_storage.go`)
- **Responsibility**: Concrete implementation using JSON files
- **Key Features**:
  - Error handling with context
  - File existence checking
  - Pretty-printed JSON output

### 3. Business Logic Layer (`manager/`)

**Purpose**: Core application logic and orchestration

#### PlaylistManager (`manager/playlist_manager.go`)
- **Responsibility**: Manage all playlists and operations
- **Key Features**:
  - Thread-safe operations using `sync.RWMutex`
  - Concurrent search with goroutines
  - CRUD operations for playlists
  - Statistics aggregation
- **Design Patterns**: 
  - Repository Pattern
  - Facade Pattern
  - Producer-Consumer (channels)

**Concurrency Model**:
```go
func (pm *PlaylistManager) SearchSongs(query string) {
    resultChan := make(chan *SearchResult, 100)
    var wg sync.WaitGroup
    
    // Spawn goroutine per playlist
    for _, playlist := range pm.playlists {
        wg.Add(1)
        go func(p *models.Playlist) {
            defer wg.Done()
            // Search logic
            resultChan <- result
        }(playlist)
    }
    
    // Wait and close
    go func() {
        wg.Wait()
        close(resultChan)
    }()
    
    // Collect results
    for result := range resultChan {
        results = append(results, result)
    }
}
```

### 4. CLI Layer (`cli/`)

**Purpose**: User interface and interaction

#### CLI (`cli/cli.go`)
- **Responsibility**: Handle user input/output
- **Key Features**:
  - Menu-driven interface
  - Input validation
  - Error display
  - Formatted output
- **Design Pattern**: Command Pattern (implicit)

## Data Flow

### Creating a Playlist and Adding a Song

```
User Input
    │
    ▼
[CLI.createPlaylist()]
    │
    ▼
[Manager.CreatePlaylist(name, description)]
    │
    ├─ Lock mutex (thread-safe)
    ├─ Create new Playlist instance
    ├─ Append to playlists slice
    └─ Unlock mutex
    │
    ▼
[Models.NewPlaylist()]
    │
    ├─ Generate unique ID
    ├─ Set timestamps
    └─ Return *Playlist
    │
    ▼
Return to CLI
    │
    ▼
Display success message
```

### Searching Songs (Concurrent)

```
User enters search query
    │
    ▼
[CLI.searchSongs()]
    │
    ▼
[Manager.SearchSongs(query)]
    │
    ├─ Acquire read lock
    ├─ Create result channel
    ├─ Create WaitGroup
    │
    ├─ For each playlist:
    │   ├─ Spawn goroutine
    │   │   ├─ Search songs
    │   │   └─ Send results to channel
    │   └─ Continue to next
    │
    ├─ Wait for all goroutines
    ├─ Close channel
    ├─ Collect all results
    └─ Release read lock
    │
    ▼
Return results to CLI
    │
    ▼
Display formatted results
```

### Saving Data

```
User exits application
    │
    ▼
[CLI.exit()]
    │
    ▼
[Manager.Save()]
    │
    ├─ Acquire read lock
    ├─ Call storage.SavePlaylists()
    │   │
    │   ▼
    │ [JSONStorage.SavePlaylists()]
    │   │
    │   ├─ Marshal to JSON
    │   ├─ Write to file
    │   └─ Return error/nil
    │
    └─ Release read lock
    │
    ▼
Display save status
    │
    ▼
Exit program
```

## Concurrency Design

### Thread Safety Strategy

1. **Read-Write Mutex** (`sync.RWMutex`):
   - Multiple readers can access simultaneously
   - Writers have exclusive access
   - Prevents race conditions

2. **Goroutines**:
   - Used for parallel search operations
   - Each playlist searched independently
   - Results communicated via channels

3. **Channels**:
   - Buffered channel for search results
   - Type-safe communication
   - Automatic synchronization

4. **Wait Groups**:
   - Track active goroutines
   - Ensure all searches complete
   - Safe channel closure

### Synchronization Points

```
┌─────────────────┐
│   Read Lock     │ ← Multiple concurrent reads OK
│  (RLock/RUnlock)│
└─────────────────┘
        │
        ├─ GetPlaylist()
        ├─ ListPlaylists()
        ├─ SearchSongs()
        └─ GetStatistics()

┌─────────────────┐
│   Write Lock    │ ← Exclusive access required
│ (Lock/Unlock)   │
└─────────────────┘
        │
        ├─ CreatePlaylist()
        ├─ DeletePlaylist()
        └─ (Playlist modifications)
```

## Design Principles Applied

### 1. **SOLID Principles**

- **Single Responsibility**: Each package has one clear purpose
- **Open/Closed**: Storage interface allows extension without modification
- **Liskov Substitution**: Any Storage implementation works
- **Interface Segregation**: Small, focused interfaces
- **Dependency Inversion**: Manager depends on Storage interface, not concrete implementation

### 2. **DRY (Don't Repeat Yourself)**

- Common formatting functions shared
- ID generation centralized
- Error handling patterns consistent

### 3. **Separation of Concerns**

- UI logic separate from business logic
- Business logic separate from data access
- Data models isolated from operations

### 4. **Composition Over Inheritance**

Go doesn't have classical inheritance. Instead:
- Struct embedding for code reuse
- Interface composition
- Method receivers for behavior

## Error Handling Strategy

### Error Propagation

```go
// Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to marshal playlists: %w", err)
}
```

### Graceful Degradation

```go
// Continue with empty state if file missing
if _, err := os.Stat(js.filepath); os.IsNotExist(err) {
    return make([]*models.Playlist, 0), nil
}
```

### User-Friendly Messages

```go
// CLI displays readable errors
if err != nil {
    fmt.Printf("❌ %v\n", err)
    return
}
```

## Testing Strategy

### Unit Testing (Suggested)

```go
// Test storage mock
type MockStorage struct {
    saveCalled bool
    playlists  []*models.Playlist
}

func (m *MockStorage) SavePlaylists(p []*models.Playlist) error {
    m.saveCalled = true
    m.playlists = p
    return nil
}

// Test manager with mock
func TestCreatePlaylist(t *testing.T) {
    mock := &MockStorage{}
    mgr := manager.NewPlaylistManager(mock)
    
    playlist := mgr.CreatePlaylist("Test", "Description")
    
    if playlist.Name != "Test" {
        t.Errorf("Expected 'Test', got '%s'", playlist.Name)
    }
}
```

### Integration Testing (Suggested)

```go
func TestFullWorkflow(t *testing.T) {
    // Create temp file
    tmpfile := "/tmp/test_playlists.json"
    defer os.Remove(tmpfile)
    
    // Create manager with real storage
    storage := storage.NewJSONStorage(tmpfile)
    mgr := manager.NewPlaylistManager(storage)
    
    // Create playlist
    p := mgr.CreatePlaylist("Test", "Description")
    
    // Add song
    song := models.NewSong("Title", "Artist", "Album", 
                          3*time.Minute, "Rock", 2023)
    p.AddSong(song)
    
    // Save
    err := mgr.Save()
    if err != nil {
        t.Fatal(err)
    }
    
    // Load in new manager
    mgr2 := manager.NewPlaylistManager(storage)
    err = mgr2.Load()
    if err != nil {
        t.Fatal(err)
    }
    
    // Verify
    playlists := mgr2.ListPlaylists()
    if len(playlists) != 1 {
        t.Errorf("Expected 1 playlist, got %d", len(playlists))
    }
}
```

## Performance Considerations

### Time Complexity

| Operation | Complexity | Notes |
|-----------|------------|-------|
| Create Playlist | O(1) | Append to slice |
| Get Playlist | O(n) | Linear search by ID |
| List Playlists | O(n) | Copy all playlists |
| Add Song | O(1) | Append to slice |
| Remove Song | O(n) | Find and remove |
| Search Songs | O(n×m) | n playlists, m songs (parallelized) |
| Shuffle | O(n) | Fisher-Yates shuffle |

### Space Complexity

- **In-Memory**: O(n) where n is total songs across all playlists
- **On-Disk**: JSON file size proportional to data

### Optimization Opportunities

1. **Index by ID**: Use map[string]*Playlist for O(1) lookups
2. **Inverted Index**: Create genre/artist indexes for faster search
3. **Lazy Loading**: Load playlists on-demand
4. **Pagination**: For large collections
5. **Caching**: Cache search results

## Extension Points

### Adding Database Support

```go
type PostgresStorage struct {
    db *sql.DB
}

func (ps *PostgresStorage) SavePlaylists(playlists []*models.Playlist) error {
    // SQL implementation
}

func (ps *PostgresStorage) LoadPlaylists() ([]*models.Playlist, error) {
    // SQL implementation
}

// Usage
store := &PostgresStorage{db: conn}
mgr := manager.NewPlaylistManager(store)
```

### Adding REST API

```go
type APIServer struct {
    manager *manager.PlaylistManager
}

func (s *APIServer) handleGetPlaylists(w http.ResponseWriter, r *http.Request) {
    playlists := s.manager.ListPlaylists()
    json.NewEncoder(w).Encode(playlists)
}

// Routes
http.HandleFunc("/api/playlists", server.handleGetPlaylists)
http.HandleFunc("/api/playlists/{id}", server.handleGetPlaylist)
```

## Conclusion

This architecture demonstrates:
- ✅ Clean separation of concerns
- ✅ Interface-based design
- ✅ Concurrent programming
- ✅ Thread safety
- ✅ Extensibility
- ✅ Testability
- ✅ Go idioms and best practices

Perfect for showcasing programming language concepts in CS4080!

