# ⚡ Quick Start Guide

## 🚀 Get Running in 5 Minutes

### Step 1: Install Go (if not installed)

**Download**: https://go.dev/dl/

**Windows**: Download and run the `.msi` installer

**Verify**:
```powershell
go version
```

### Step 2: Run the Application

```powershell
cd C:\Users\myLAPPY\MusicPlaylistProject
go run main.go
```

**That's it!** The application will start immediately.

---

## 🎯 First Time User Guide

### Create Your First Playlist (30 seconds)

1. **Start the app**: `go run main.go`

2. **Press 1** → Create New Playlist
   - Name: `My Favorites`
   - Description: `Best songs ever`

3. **Press 4** → Add Song
   - Enter the playlist ID shown
   - Title: `Bohemian Rhapsody`
   - Artist: `Queen`
   - Album: `A Night at the Opera`
   - Duration: `5:55`
   - Genre: `Rock`
   - Year: `1975`

4. **Press 3** → View Playlist
   - See your song!

5. **Press 0** → Exit (saves automatically)

6. **Run again** → Your playlist is still there! ✨

---

## 📋 Menu Quick Reference

| Key | Action |
|-----|--------|
| 1 | Create Playlist |
| 2 | List All Playlists |
| 3 | View Playlist Details |
| 4 | Add Song |
| 5 | Remove Song |
| 6 | Search Songs |
| 7 | Shuffle Playlist |
| 8 | Delete Playlist |
| 9 | Show Statistics |
| 0 | Save & Exit |

---

## 🎵 Sample Data to Try

### Rock Classics Playlist

```
Song 1:
- Title: Stairway to Heaven
- Artist: Led Zeppelin
- Album: Led Zeppelin IV
- Duration: 8:02
- Genre: Rock
- Year: 1971

Song 2:
- Title: Hotel California
- Artist: Eagles
- Album: Hotel California
- Duration: 6:30
- Genre: Rock
- Year: 1977

Song 3:
- Title: Comfortably Numb
- Artist: Pink Floyd
- Album: The Wall
- Duration: 6:23
- Genre: Progressive Rock
- Year: 1979
```

### 80s Pop Playlist

```
Song 1:
- Title: Billie Jean
- Artist: Michael Jackson
- Album: Thriller
- Duration: 4:54
- Genre: Pop
- Year: 1983

Song 2:
- Title: Like a Prayer
- Artist: Madonna
- Album: Like a Prayer
- Duration: 5:43
- Genre: Pop
- Year: 1989

Song 3:
- Title: Sweet Child O' Mine
- Artist: Guns N' Roses
- Album: Appetite for Destruction
- Duration: 5:56
- Genre: Rock
- Year: 1987
```

---

## 🔍 Testing Features

### Test Concurrent Search
1. Create 3-4 playlists with several songs each
2. Press **6** (Search)
3. Search for "rock" or an artist name
4. Watch it search across all playlists instantly! ⚡

### Test Shuffle
1. Create a playlist with 5+ songs
2. View it (note the order)
3. Press **7** (Shuffle)
4. View it again (order changed!) 🔀

### Test Persistence
1. Add playlists and songs
2. Press **0** (Exit)
3. Run the app again
4. All your data is still there! 💾

### Test Statistics
1. Create playlists with different genres
2. Press **9** (Statistics)
3. See genre breakdown and artist counts 📊

---

## 🏗️ Build Executable (Optional)

```powershell
go build -o playlist-manager.exe
.\playlist-manager.exe
```

Now you have a standalone executable!

---

## 📁 Where is My Data?

Your playlists are saved in:
```
C:\Users\myLAPPY\MusicPlaylistProject\playlists.json
```

**Backup**: Just copy this file!  
**Share**: Send this file to a friend!  
**Reset**: Delete this file to start fresh!

---

## 🐛 Common Issues

### "go: command not found"
**Solution**: Install Go or restart your terminal

### "module not found"
**Solution**: 
```powershell
go mod tidy
```

### "permission denied"
**Solution**: Run PowerShell as Administrator

---

## 🎓 For CS4080 Demo

### Quick Demo Script (5 minutes)

```
1. Start app
   ⏱️ "This is a music playlist manager built in Go"

2. Create playlist (1)
   ⏱️ "It allows creating organized collections"

3. Add 2-3 songs (4)
   ⏱️ "Each song has rich metadata"

4. Search (6)
   ⏱️ "Search uses goroutines for concurrent processing"

5. Statistics (9)
   ⏱️ "Aggregates data across all playlists"

6. Exit and restart (0)
   ⏱️ "Data persists between sessions using JSON"

7. Show code
   ⏱️ "Clean architecture with interfaces, concurrency, and thread safety"
```

### Key Points to Mention

✅ **Interfaces** for abstraction (Storage)  
✅ **Goroutines** for concurrency (Search)  
✅ **Channels** for communication  
✅ **Mutex locks** for thread safety  
✅ **Error handling** explicit returns  
✅ **Composition** over inheritance  

---

## 📚 Documentation Files

- **README.md** → Project overview
- **INSTALLATION.md** → Setup instructions
- **EXAMPLES.md** → Usage examples
- **ARCHITECTURE.md** → Technical design
- **PROJECT_SUMMARY.md** → CS4080 concepts
- **QUICK_START.md** → This file!

---

## 🚀 Next Steps

### Extend the Project

1. **Add Ratings**: Let users rate songs 1-5 stars
2. **Play History**: Track last played time
3. **Export Playlist**: Save to M3U format
4. **Import CSV**: Bulk import songs
5. **Web Interface**: Add HTTP server
6. **Database**: Use SQLite or PostgreSQL

### Learn More Go

- Take the **Go Tour**: https://go.dev/tour/
- Read **Effective Go**: https://go.dev/doc/effective_go
- Explore **Go by Example**: https://gobyexample.com/

---

## 🎉 You're Ready!

Your music playlist manager is:
- ✅ Complete and working
- ✅ Well-documented
- ✅ Demonstrates key PL concepts
- ✅ Ready for CS4080 submission

**Now go create some playlists! 🎵**

---

**Pro Tip**: Copy `QUICK_START.md` to show your professor how easy it is to use! 😎

