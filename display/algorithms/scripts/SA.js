// Load data based on selected run
async function loadData(runNumber) {
    try {
        const response = await fetch(`/cubes/SA/Simulated Annealing_${runNumber}.json`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        data = await response.json();
        
        // Update info
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('InitT').textContent = data.initialTemp;
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states[data.states.length-1].iteration;
        document.getElementById('stuck').textContent = 1 - (data.customVar / (data.states[data.states.length-1].iteration));
        document.getElementById('Duration').textContent = formatDuration(data.duration);

        // Create chart data - Include prob value
        const chartData = data.states.map(state => ({
            iteration: state.iteration,
            OF: state.OF,
            prob: state.prob  // Add this!
        }));

        // Create both charts
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

// Display state information and cube
function displayState(stateIndex) {
    // Remove iteration check as we're using stateIndex directly
    const state = data.states[stateIndex];
    
    // Update state information
    document.getElementById('iteration').textContent = state.iteration;
    document.getElementById('ofValue').textContent = state.OF;
    document.getElementById('temp').textContent = state.temperature;
    document.getElementById('prob').textContent = state.prob ? state.prob.toFixed(4) : '-'; // Format probability to 4 decimal places
    document.getElementById('action').textContent = state.action;
    document.getElementById('progress').value = stateIndex;
    
    // Display cube
    displayCube(state.cube);

    if (myChart) {
        currentStateIndex = stateIndex;  
        myChart.update();
    }
    if (probChart) {
        currentStateIndex = stateIndex;
        probChart.update();
    }
}