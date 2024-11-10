// Common variables and functions that might be used across different algorithms
let data = null;
let isPlaying = false;
let currentStateIndex = 0;
let playbackSpeed = 1;
let rotationX = -30;
let rotationY = 45;

// Format duration from nanoseconds to readable format
function formatDuration(nanoseconds) {
    const milliseconds = nanoseconds / 1000000;
    if (milliseconds < 1000) {
        return `${milliseconds.toFixed(2)}ms`;
    }
    const seconds = milliseconds / 1000;
    return `${seconds.toFixed(2)}s`;
}

// Event listeners
document.addEventListener('DOMContentLoaded', function() {
    // Hide temperature element
    const tempElement = document.getElementById('temperature');
    if (tempElement && tempElement.parentElement) {
        tempElement.parentElement.style.display = 'none';
    }
    
    // Run selector
    document.getElementById('runSelect').addEventListener('change', function(e) {
        const runNumber = e.target.value.replace('run', '');
        loadData(runNumber);
    });

    // Play/Pause button
    document.getElementById('playPause').addEventListener('click', togglePlay);

    // Progress bar
    document.getElementById('progress').addEventListener('input', function(e) {
        currentStateIndex = parseInt(e.target.value);
        displayState(currentStateIndex);
        if (isPlaying) {
            clearInterval(playbackInterval);
            togglePlay();
        }
    });

    // Speed selector
    document.getElementById('speed').addEventListener('change', function() {
        if (isPlaying) {
            clearInterval(playbackInterval);
            togglePlay();
        }
    });

    // Initial load
    loadData(1);
});