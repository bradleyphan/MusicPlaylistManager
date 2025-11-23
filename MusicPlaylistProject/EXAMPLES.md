# 🎵 Usage Examples

This document provides practical examples of using the Music Playlist Manager.

## Example 1: Creating Your First Playlist

```
Main Menu -> Select 1 (Create New Playlist)
Playlist name: Road Trip 2024
Description: Best songs for long drives
✅ Created playlist: [P1732234567890] Road Trip 2024 - 0 songs (0s total)
```

## Example 2: Adding Songs

```
Main Menu -> Select 4 (Add Song to Playlist)
Enter playlist ID: P1732234567890

ADD NEW SONG
Title: Bohemian Rhapsody
Artist: Queen
Album: A Night at the Opera
Genre: Rock
Duration (MM:SS): 5:55
Year: 1975
✅ Added song: [S1732234567891] Bohemian Rhapsody - Queen (A Night at the Opera) [Rock] - 5:55
```

## Example 3: Building a Complete Playlist

Here's a sample playlist you can recreate:

**Playlist: "90s Hits"**

| Title | Artist | Album | Duration | Genre | Year |
|-------|--------|-------|----------|-------|------|
| Smells Like Teen Spirit | Nirvana | Nevermind | 5:01 | Grunge | 1991 |
| Wonderwall | Oasis | What's the Story Morning Glory | 4:18 | Rock | 1995 |
| No Scrubs | TLC | FanMail | 3:35 | R&B | 1999 |
| Wannabe | Spice Girls | Spice | 2:53 | Pop | 1996 |
| Black Hole Sun | Soundgarden | Superunknown | 5:18 | Rock | 1994 |

## Example 4: Searching Songs

```
Main Menu -> Select 6 (Search Songs)
🔍 Enter search query: queen

✅ Found 2 result(s):
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. [S1732234567891] Bohemian Rhapsody - Queen (A Night at the Opera) [Rock] - 5:55 (in playlist: Road Trip 2024)
2. [S1732234567892] We Will Rock You - Queen (News of the World) [Rock] - 2:02 (in playlist: Classic Rock)
```

## Example 5: Viewing Statistics

```
Main Menu -> Select 9 (Show Statistics)

📊 STATISTICS
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Total Playlists: 3
Total Songs: 15
Total Duration: 1h 2m 45s

🎸 Top Genres:
  - Rock: 8 songs
  - Pop: 4 songs
  - R&B: 2 songs
  - Grunge: 1 songs

🎤 Top Artists:
  - Queen: 3 songs
  - Nirvana: 2 songs
  - Oasis: 2 songs
  - TLC: 1 songs
```

## Example 6: Shuffling a Playlist

```
Before Shuffle:
1. [S001] Song A
2. [S002] Song B
3. [S003] Song C
4. [S004] Song D

Main Menu -> Select 7 (Shuffle Playlist)
Enter playlist ID: P1732234567890
✅ Playlist shuffled successfully!

After Shuffle:
1. [S003] Song C
2. [S001] Song A
3. [S004] Song D
4. [S002] Song B
```

## Example 7: Complete Session

```
🎵 Music Playlist Manager - Go Edition
CS4080 Programming Languages Project

Loading playlists...
Starting with empty playlist collection.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📋 MAIN MENU
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. Create New Playlist
2. List All Playlists
...

Enter your choice: 1

➕ CREATE NEW PLAYLIST
Playlist name: Workout Mix
Description: High energy songs for gym
✅ Created playlist: [P1732234567890] Workout Mix - 0 songs (0s total)

Enter your choice: 4

📚 ALL PLAYLISTS:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. [P1732234567890] Workout Mix - 0 songs (0s total)

Enter playlist ID: P1732234567890

🎵 ADD NEW SONG
Title: Eye of the Tiger
Artist: Survivor
Album: Eye of the Tiger
Genre: Rock
Duration (MM:SS): 4:05
Year: 1982
✅ Added song: [S1732234567891] Eye of the Tiger - Survivor (Eye of the Tiger) [Rock] - 4:05

Enter your choice: 2

📚 ALL PLAYLISTS:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
1. [P1732234567890] Workout Mix - 1 songs (4:05 total)

Enter your choice: 0

💾 Saving data...
✅ Data saved successfully!
👋 Goodbye!
```

## Tips & Tricks

### 1. Quick Duration Entry
- Use MM:SS format: `3:45` for 3 minutes 45 seconds
- For hours: multiply by 60: `75:30` for 1 hour 15 minutes 30 seconds

### 2. Search Tips
- Search is case-insensitive
- Searches across: title, artist, album, and genre
- Partial matches work: "queen" finds "Queen" and "Queensrÿche"

### 3. Organizing Playlists
- Use descriptive names: "Morning Coffee" not "Playlist1"
- Add detailed descriptions to remember the purpose
- Create theme-based playlists for easier management

### 4. Backup Your Data
- The `playlists.json` file contains all your data
- Copy it regularly to backup your playlists
- You can share it with others or move it between computers

### 5. Performance
- Search uses concurrent goroutines for speed
- Handles thousands of songs efficiently
- All operations are thread-safe

## Common Use Cases

### Workout Playlist
Create high-energy playlists for different workout types:
- **Cardio**: 140-180 BPM songs
- **Weightlifting**: Heavy rock/metal
- **Yoga**: Calm, ambient music

### Study Playlist
Organize focus music:
- **Deep Focus**: Instrumental, ambient
- **Light Background**: Soft vocals, jazz
- **Break Time**: Upbeat favorites

### Party Playlist
Build the perfect party mix:
- **Warm-up**: Easy listening, classics
- **Peak Hours**: Dance, pop hits
- **Wind Down**: Slower tempo favorites

## Programming Concepts in Action

### Concurrent Search
When you search for songs, the system creates a separate goroutine for each playlist, allowing parallel searching. For 10 playlists with 100 songs each, this provides significant speedup.

### Thread Safety
Multiple operations can safely run concurrently thanks to read-write mutex locks. Try opening multiple terminal windows and running operations simultaneously!

### Data Persistence
The JSON storage demonstrates separation of concerns - the manager doesn't know *how* data is saved, just that it implements the Storage interface. This makes it easy to swap JSON for a database later.

---

**Enjoy managing your playlists! 🎵**

