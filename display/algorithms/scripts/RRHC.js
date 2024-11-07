// Load data based on selected run
async function loadData(runNumber) {
    try {
        console.log('Attempting to load run', runNumber);
        const response = await fetch(`/cubes/RRHC/Random Restart Hill Climb_${runNumber}.json`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        data = await response.json();
        console.log('Loaded data:', data);
        
        // Update UI with loaded data
        document.getElementById('Restart').textContent = data.forRRHC.length;
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states.length;
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        
        // Display initial state
        displayState(0);
        
        return true;
    } catch (error) {
        console.error('Error loading data:', error);
        return false;
    }
}

// Display state information
function displayState(stateIndex) {
    const state = data.states[stateIndex];
    document.getElementById('iteration').textContent = state.iteration;
    document.getElementById('ofValue').textContent = state.OF;
    document.getElementById('action').textContent = state.action;
    // Hide temperature since it's not used in SAHC
    document.getElementById('temperature').parentElement.style.display = 'none';
}

// Event listeners
document.addEventListener('DOMContentLoaded', function() {
    // Hide temperature element since SAHC doesn't use it
    document.getElementById('temperature').parentElement.style.display = 'none';
    
    // Run selector
    document.getElementById('runSelect').addEventListener('change', function(e) {
        const runNumber = e.target.value.replace('run', '');
        loadData(runNumber);
    });

    // Initial load
    loadData(1);
});