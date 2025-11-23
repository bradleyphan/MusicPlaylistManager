# 🚀 Quick Run Guide

## Choose Your Interface

### 🖥️ Command-Line Interface (CLI)

**Best for:** Code demonstrations, showing Go concepts

```powershell
go run main.go
```

Then follow the interactive menu (press 1, 2, 3, etc.)

---

### 🌐 Web Interface (RECOMMENDED!)

**Best for:** Presentations, visual demos, impressing your professor!

```powershell
go run main_web.go
```

Then open your browser to:
```
http://localhost:8080
```

**Features:**
- ✨ Beautiful modern design
- 🎨 Gradient backgrounds
- 📊 Real-time statistics dashboard
- 🖱️ Point-and-click interface
- 📱 Works on phone/tablet too!

---

## First Time Setup

1. **Install Go** (if not already installed)
   - Download from: https://go.dev/dl/
   - Windows: Run the .msi installer
   - Verify: `go version`

2. **Navigate to project**
   ```powershell
   cd C:\Users\myLAPPY\MusicPlaylistProject
   ```

3. **Run either version** (see above)

---

## Sample Data to Try

### Create a Playlist
- Name: "Road Trip 2024"
- Description: "Best driving songs"

### Add Some Songs
1. **Bohemian Rhapsody**
   - Artist: Queen
   - Album: A Night at the Opera
   - Duration: 355 seconds (CLI) or just type in web
   - Genre: Rock
   - Year: 1975

2. **Hotel California**
   - Artist: Eagles
   - Album: Hotel California
   - Duration: 390 seconds
   - Genre: Rock
   - Year: 1977

3. **Billie Jean**
   - Artist: Michael Jackson
   - Album: Thriller
   - Duration: 294 seconds
   - Genre: Pop
   - Year: 1983

---

## 🎓 For Your CS4080 Presentation

### Demo Strategy

**1. Start with Web Interface (3 min)**
- Show the beautiful UI
- Create a playlist live
- Add a few songs
- Show search feature
- Demonstrate shuffle

**2. Show the CLI (2 min)**
- Exit web server
- Run `go run main.go`
- Show same data is there!
- Demonstrate text interface

**3. Explain the Code (5 min)**
- Open `manager/playlist_manager.go`
- Show concurrent search with goroutines
- Explain interfaces (`storage/storage.go`)
- Point out thread safety (mutex locks)
- Show how both interfaces share the backend

**4. Q&A (2 min)**

---

## Quick Comparison

| Feature | CLI | Web |
|---------|-----|-----|
| **Visual Appeal** | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Ease of Use** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **Code Demo** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |
| **Wow Factor** | ⭐⭐ | ⭐⭐⭐⭐⭐ |

**Recommendation:** Start with web for visual impact, then show CLI to explain code!

---

## Troubleshooting

### "go: command not found"
- Restart your terminal after installing Go
- Or run: 
  ```powershell
  $env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User")
  ```

### Web server port already in use
- Another program is using port 8080
- Edit `main_web.go` and change `:8080` to `:3000`

### Changes not appearing
- Web UI auto-refreshes every 5 seconds
- Or manually refresh your browser (F5)

---

## 📚 More Documentation

- **README.md** - Complete project overview
- **WEB_GUIDE.md** - Detailed web interface guide
- **QUICK_START.md** - 5-minute tutorial
- **ARCHITECTURE.md** - Technical design details
- **PROJECT_SUMMARY.md** - CS4080 concepts explained

---

**You got this! 🎉**

