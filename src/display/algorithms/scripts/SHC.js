// Load data based on selected run
async function loadData(runNumber) {
    try {
        const response = await fetch(`/cubes/SHC/Stochastic Hill Climb_${runNumber}.json`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        data = await response.json();
        console.log('Loaded data:', data);
        
        // Update info
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('MaxI').textContent = data.customVar;
        document.getElementById('TotalI').textContent = data.states[data.states.length-1].iteration;
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        document.getElementById('Comp').textContent = (100 - ((data.finalOF / data.initialOF)*100)).toFixed(4);

        // Create chart data
        const chartData = data.states.map(state => ({
            iteration: state.iteration,
            OF: state.OF
        }));

        // Create chart
        createChart(chartData);
        
        // Initialize player controls with iterations instead of state indices
        document.getElementById('progress').max = data.states[data.states.length-1].iteration;
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