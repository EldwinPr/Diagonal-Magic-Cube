// Common variables and functions that might be used across different algorithms
let data = null;
let isPlaying = false;
let currentStateIndex = 0;
let playbackSpeed = 1;

// Format duration from nanoseconds to readable format
function formatDuration(nanoseconds) {
    const milliseconds = nanoseconds / 1000000;
    if (milliseconds < 1000) {
        return `${milliseconds.toFixed(2)}ms`;
    }
    const seconds = milliseconds / 1000;
    return `${seconds.toFixed(2)}s`;
}