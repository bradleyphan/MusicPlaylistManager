# 📦 Installation Guide

## Installing Go on Windows

### Method 1: Official Installer (Recommended)

1. **Download Go**:
   - Visit https://go.dev/dl/
   - Download the Windows installer (`.msi` file)
   - Current recommended version: Go 1.21 or higher

2. **Run the Installer**:
   - Double-click the downloaded `.msi` file
   - Follow the installation wizard
   - Default installation path: `C:\Program Files\Go`

3. **Verify Installation**:
   ```powershell
   go version
   ```
   You should see something like: `go version go1.21.x windows/amd64`

4. **Set Up Go Workspace** (Usually automatic):
   ```powershell
   go env GOPATH
   ```
   Default: `C:\Users\YourUsername\go`

### Method 2: Using Chocolatey

If you have Chocolatey installed:

```powershell
choco install golang
```

### Method 3: Using Scoop

If you have Scoop installed:

```powershell
scoop install go
```

## Setting Up the Project

1. **Navigate to Project Directory**:
   ```powershell
   cd C:\Users\myLAPPY\MusicPlaylistProject
   ```

2. **Initialize Go Modules** (Already done):
   ```powershell
   go mod tidy
   ```

3. **Build the Application**:
   ```powershell
   go build -o playlist-manager.exe
   ```

4. **Run the Application**:
   ```powershell
   .\playlist-manager.exe
   ```

   Or run directly without building:
   ```powershell
   go run main.go
   ```

## Troubleshooting

### "go: command not found" or "not recognized"

**Solution**: Restart your terminal after installing Go. If still not working:

1. Open Environment Variables:
   - Press `Win + X` → System → Advanced system settings
   - Click "Environment Variables"

2. Check PATH includes:
   - `C:\Program Files\Go\bin`
   - `%USERPROFILE%\go\bin`

3. Restart PowerShell/CMD

### Module Import Errors

If you see import errors:

```powershell
go mod init musicplaylist
go mod tidy
```

### Permission Errors

Run PowerShell as Administrator if you encounter permission issues.

## Recommended Development Tools

### 1. **Visual Studio Code**
- Download: https://code.visualstudio.com/
- Install Go extension by Go Team at Google
- Features: IntelliSense, debugging, linting

### 2. **GoLand** (JetBrains)
- Download: https://www.jetbrains.com/go/
- Full-featured IDE for Go
- Free for students with JetBrains Student License

### 3. **Git** (for version control)
- Download: https://git-scm.com/download/win
- Initialize repository:
  ```powershell
  git init
  git add .
  git commit -m "Initial commit: Music Playlist Manager"
  ```

## IDE Configuration

### Visual Studio Code

1. Install Go extension
2. Open project folder
3. Press `Ctrl+Shift+P` → "Go: Install/Update Tools"
4. Select all tools and install

**Recommended settings** (`.vscode/settings.json`):
```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.formatTool": "goimports",
  "editor.formatOnSave": true
}
```

## Building for Different Platforms

### Windows (64-bit)
```powershell
go build -o playlist-manager.exe
```

### Windows (32-bit)
```powershell
$env:GOARCH="386"
go build -o playlist-manager-x86.exe
```

### Linux
```powershell
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o playlist-manager
```

### macOS
```powershell
$env:GOOS="darwin"
$env:GOARCH="amd64"
go build -o playlist-manager
```

## Running Tests (Future)

When you add tests:

```powershell
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Run with verbose output
go test -v ./...

# Run specific package
go test ./models
```

## Common Go Commands

```powershell
# Format code
go fmt ./...

# Check for errors
go vet ./...

# Download dependencies
go mod download

# Clean build cache
go clean

# View documentation
go doc models.Song

# Install as global command
go install
```

## Verifying Project Structure

Your project should look like this:

```
MusicPlaylistProject/
├── go.mod                  ✓
├── main.go                 ✓
├── models/
│   ├── song.go            ✓
│   └── playlist.go        ✓
├── storage/
│   ├── storage.go         ✓
│   └── json_storage.go    ✓
├── manager/
│   └── playlist_manager.go ✓
├── cli/
│   └── cli.go             ✓
├── README.md              ✓
├── EXAMPLES.md            ✓
├── ARCHITECTURE.md        ✓
├── INSTALLATION.md        ✓
└── .gitignore             ✓
```

## First Run Checklist

- [ ] Go is installed (`go version` works)
- [ ] Navigated to project directory
- [ ] Built the project successfully
- [ ] Ran the application
- [ ] Created a test playlist
- [ ] Added a test song
- [ ] Exited and verified data persistence

## Need Help?

### Official Resources
- Go Documentation: https://go.dev/doc/
- Go Tour (Interactive): https://go.dev/tour/
- Effective Go: https://go.dev/doc/effective_go

### Community
- Go Forum: https://forum.golangbridge.org/
- Reddit: r/golang
- Stack Overflow: Tag `go`

## For CS4080 Submission

### What to Submit

1. **Source Code**: Entire `MusicPlaylistProject` folder
2. **Documentation**: All `.md` files included
3. **Demonstration**: Screenshots or video of running application
4. **Report**: Explain programming concepts demonstrated (use ARCHITECTURE.md as reference)

### Recommended Demo Steps

1. Show the application running
2. Create 2-3 playlists
3. Add multiple songs
4. Demonstrate search functionality
5. Show statistics
6. Exit and restart to show persistence
7. Walk through code highlighting:
   - Interface design (Storage)
   - Concurrency (SearchSongs)
   - Thread safety (mutex locks)
   - Error handling

---

**Good luck with your CS4080 project! 🚀**

