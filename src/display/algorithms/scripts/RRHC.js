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
        
        // Update info
        document.getElementById('Restart').textContent = data.customArr.length;
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('FinalOF').textContent = data.states[data.states.length - 1].OF;
        document.getElementById('BestOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states.length - 1;
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        document.getElementById('Comp').textContent = (100 - ((data.finalOF / data.initialOF)*100)).toFixed(4);
        
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

    // find restart number
    let resnum = 1;
    for (let i = 0; i < data.customArr.length; i++) {
        if (state.iteration >= data.customArr[i]) {
            resnum++;
        }
    }

    // Update state information
    document.getElementById('RestNum').textContent = resnum;
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