// Load data based on selected run
async function loadData(runNumber) {
    try {
        console.log('Attempting to load run', runNumber);
        const response = await fetch(`/cubes/HCWSM/Hill Climb with Sideways Moves_${runNumber}.json`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        data = await response.json();
        console.log('Loaded data:', data);
        
        // Update info
        document.getElementById('MaxSM').textContent = data.customVar;
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states.length;
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        
        // Create chart data
        const chartData = data.states.map(state => ({
            iteration: state.iteration,
            OF: state.OF
        }));

        // Create chart
        createChart(chartData);

        // Initialize player controls
        document.getElementById('progress').max = data.states.length - 1;
        document.getElementById('progress').value = 0;
        currentStateIndex = 0;
        isPlaying = false;
        document.getElementById('playPause').textContent = 'Play';
        
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

    // Update state information
    document.getElementById('iteration').textContent = state.iteration;
    document.getElementById('ofValue').textContent = state.OF;
    document.getElementById('action').textContent = state.action;
    document.getElementById('progress').value = stateIndex;

    // Display cube
    displayCube(state.cube);

    if (myChart) {
        myChart.update();
    }
}