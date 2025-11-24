// Global state
let currentPlaylistId = null;
let playlists = [];

document.addEventListener('DOMContentLoaded', function() {
    loadPlaylists();
    loadStatistics();
    
    // Refresh every 5 seconds
    setInterval(() => {
        loadPlaylists();
        loadStatistics();
    }, 5000);
});

async function loadPlaylists() {
    try {
        const response = await fetch('/api/playlists');
        playlists = await response.json();
        renderPlaylists();
        
        // Reload current playlist if selected
        if (currentPlaylistId) {
            const playlist = playlists.find(p => p.id === currentPlaylistId);
            if (playlist) {
                renderSongs(playlist);
            }
        }
    } catch (error) {
        console.error('Error loading playlists:', error);
    }
}

async function loadStatistics() {
    try {
        const response = await fetch('/api/statistics');
        const stats = await response.json();
        
        document.getElementById('totalPlaylists').textContent = stats.TotalPlaylists || 0;
        document.getElementById('totalSongs').textContent = stats.TotalSongs || 0;
        
        const duration = stats.TotalDuration || 0;
        const minutes = Math.floor(duration / 60000000000);
        document.getElementById('totalDuration').textContent = minutes + 'm';
    } catch (error) {
        console.error('Error loading statistics:', error);
    }
}

function renderPlaylists() {
    const container = document.getElementById('playlistsList');
    
    if (playlists.length === 0) {
        container.innerHTML = '<div class="empty-state"><p>No playlists yet. Create one!</p></div>';
        return;
    }
    
    container.innerHTML = playlists.map(playlist => `
        <div class="playlist-item ${playlist.id === currentPlaylistId ? 'active' : ''}" 
             onclick="selectPlaylist('${playlist.id}')">
            <div class="playlist-name">${escapeHtml(playlist.name)}</div>
            <div class="playlist-info">
                ${playlist.songs ? playlist.songs.length : 0} songs • 
                ${formatDuration(calculateTotalDuration(playlist.songs))}
            </div>
        </div>
    `).join('');
}

function selectPlaylist(playlistId) {
    currentPlaylistId = playlistId;
    const playlist = playlists.find(p => p.id === playlistId);
    
    if (playlist) {
        renderSongs(playlist);
        renderPlaylists(); // Re-render to update active state
    }
}

function renderSongs(playlist) {
    const container = document.getElementById('songsList');
    const titleElement = document.getElementById('playlistTitle');
    const actionsElement = document.getElementById('playlistActions');
    
    titleElement.textContent = `🎵 ${playlist.name}`;
    actionsElement.style.display = 'flex';
    
    if (!playlist.songs || playlist.songs.length === 0) {
        container.innerHTML = '<div class="empty-state"><p>No songs in this playlist yet</p></div>';
        return;
    }
    
    container.innerHTML = playlist.songs.map(song => `
        <div class="song-item">
            <div class="song-details">
                <div class="song-title">${escapeHtml(song.title)}</div>
                <div class="song-meta">
                    ${escapeHtml(song.artist)} • ${escapeHtml(song.album)} • 
                    ${escapeHtml(song.genre)} • ${formatDuration(song.duration)}
                </div>
            </div>
            <div class="song-actions">
                <button onclick="removeSong('${song.id}')">Remove</button>
            </div>
        </div>
    `).join('');
}

async function createPlaylist(event) {
    event.preventDefault();
    
    const name = document.getElementById('playlistName').value;
    const description = document.getElementById('playlistDescription').value;
    
    try {
        const response = await fetch('/api/playlists/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, description })
        });
        
        if (response.ok) {
            closeModal('createPlaylistModal');
            document.getElementById('playlistName').value = '';
            document.getElementById('playlistDescription').value = '';
            loadPlaylists();
            loadStatistics();
        }
    } catch (error) {
        alert('Error creating playlist: ' + error.message);
    }
}

async function addSong(event) {
    event.preventDefault();
    
    if (!currentPlaylistId) {
        alert('Please select a playlist first');
        return;
    }
    
    const song = {
        playlist_id: currentPlaylistId,
        file_path: document.getElementById('filePath').value,
        duration: parseInt(document.getElementById('songDuration').value),
    };
    
    try {
        const response = await fetch('/api/songs/add', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(song)
        });
        
        if (response.ok) {
            closeModal('addSongModal');

            document.getElementById('filePath').value = '';
            loadPlaylists();
            loadStatistics();
        }else{
            alert("Error adding song")
        }
    } catch (error) {
        alert('Error adding song: ' + error.message);
    }
}
async function scanFolder(event) {
    event.preventDefault();
    
    if (!currentPlaylistId) {
        alert('Please select a playlist first');
        return;
    }
    
    const folder = {
        playlist_id: currentPlaylistId,
        file_path: document.getElementById('folderPath').value,
    };
    
    try {
        const response = await fetch('/api/songs/scan', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(folder)
        });
        
        if (response.ok) {
            closeModal('scanFolderModal');
            document.getElementById('filePath').value = '';
            loadPlaylists();
            loadStatistics();
        }else{
            alert("Error scanning folder")
        }
    } catch (error) {
        alert('Error : ' + error.message);
    }
}

async function removeSong(songId) {
    if (!confirm('Remove this song?')) return;
    
    try {
        const response = await fetch('/api/songs/remove', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                playlist_id: currentPlaylistId,
                song_id: songId
            })
        });
        
        if (response.ok) {
            loadPlaylists();
            loadStatistics();
        }
    } catch (error) {
        alert('Error removing song: ' + error.message);
    }
}

async function deletePlaylist() {
    if (!currentPlaylistId) return;
    
    if (!confirm('Delete this entire playlist? This cannot be undone!')) return;
    
    try {
        const response = await fetch('/api/playlists/delete', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id: currentPlaylistId })
        });
        
        if (response.ok) {
            currentPlaylistId = null;
            document.getElementById('songsList').innerHTML = '<div class="empty-state"><p>Select a playlist to view songs</p></div>';
            document.getElementById('playlistTitle').textContent = 'Select a playlist';
            document.getElementById('playlistActions').style.display = 'none';
            loadPlaylists();
            loadStatistics();
        }
    } catch (error) {
        alert('Error deleting playlist: ' + error.message);
    }
}

async function shufflePlaylist() {
    if (!currentPlaylistId) return;
    
    try {
        const response = await fetch('/api/playlists/shuffle', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ id: currentPlaylistId })
        });
        
        if (response.ok) {
            loadPlaylists();
        }
    } catch (error) {
        alert('Error shuffling playlist: ' + error.message);
    }
}

async function searchSongs() {
    const query = document.getElementById('searchInput').value.trim();
    
    if (query.length < 2) {
        return;
    }
    
    try {
        const response = await fetch(`/api/songs/search?q=${encodeURIComponent(query)}`);
        const results = await response.json();
        
        if (results && results.length > 0) {
            showSearchResults(results);
        }
    } catch (error) {
        console.error('Error searching:', error);
    }
}

function showSearchResults(results) {
    const modal = document.getElementById('searchResultsModal');
    const container = document.getElementById('searchResults');
    
    container.innerHTML = results.map(result => `
        <div class="search-result-item">
            <div class="search-result-song">
                ${escapeHtml(result.Song.title)} - ${escapeHtml(result.Song.artist)}
            </div>
            <div class="search-result-playlist">
                In playlist: ${escapeHtml(result.PlaylistName)}
            </div>
        </div>
    `).join('');
    
    modal.style.display = 'block';
}

// Modal functions
function showCreatePlaylistModal() {
    document.getElementById('createPlaylistModal').style.display = 'block';
}

function showAddSongModal() {
    if (!currentPlaylistId) {
        alert('Please select a playlist first');
        return;
    }
    document.getElementById('addSongModal').style.display = 'block';
}

function showScanFolderModal() {
    if (!currentPlaylistId) {
        alert('Please select a playlist first');
        return;
    }
    document.getElementById('scanFolderModal').style.display = 'block';
}

function closeModal(modalId) {
    document.getElementById(modalId).style.display = 'none';
}

// Close modal when clicking outside
window.onclick = function(event) {
    if (event.target.classList.contains('modal')) {
        event.target.style.display = 'none';
    }
}

function formatDuration(duration) {
    const seconds = Math.floor(duration / 1000000000);
    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;
    return `${minutes}:${remainingSeconds.toString().padStart(2, '0')}`;
}

function calculateTotalDuration(songs) {
    if (!songs || songs.length === 0) return 0;
    return songs.reduce((total, song) => total + (song.duration || 0), 0);
}

function escapeHtml(text) {
    if (!text) return '';
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

