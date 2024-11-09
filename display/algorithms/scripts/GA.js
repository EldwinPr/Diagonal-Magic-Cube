async function loadData(runNumber) {
    try {
        const response = await fetch(`/cubes/GA/Genetic Algorithm_${runNumber}.json`);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        data = await response.json();
        console.log('Loaded data:', data);
        
        // Update info
        document.getElementById('InitOF').textContent = data.initialOF;
        document.getElementById('Pop').textContent = data.customVar;
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states.length - 1;
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        document.getElementById('Comp').textContent = (100 - ((data.finalOF / data.initialOF)*100)).toFixed(4);

        // Create chart data
        const chartData = data.states.map((state, index) => ({
            iteration: state.iteration,
            OF: state.OF,
            AvgOF: data.customArr[index] // Map the average OF from customArr at the same index
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

// Display state information and cube
function displayState(stateIndex) {
    const state = data.states[stateIndex];
    
    // Update state information
    document.getElementById('iteration').textContent = state.iteration;
    document.getElementById('ofValue').textContent = state.OF;
    document.getElementById('action').textContent = state.action;
    document.getElementById('progress').value = stateIndex;
    
    // Update avgOF display correctly
    if (data.customArr && data.customArr[stateIndex] !== undefined) {
        document.getElementById('avgOF').textContent = data.customArr[stateIndex];
    }

    // Display cube
    displayCube(state.cube);

    if (myChart) {
        myChart.update();
    }
    if (avgChart) {
        avgChart.update();
    }
}

//Chart Creation
let myChart = null;
let avgChart = null;

function createChart(chartData) {
    const canvas = document.getElementById('ofChart');
    const probCanvas = document.getElementById('probChart');
    const avgCanvas = document.getElementById('avgChart');
    
    // Safe destroy with null checks
    if (myChart instanceof Chart) {
        myChart.destroy();
    }
    if (avgChart instanceof Chart) {
        avgChart.destroy();
    }

    // Reset to null after destroy
    myChart = null;
    avgChart = null;

    // Create OF chart
    myChart = new Chart(canvas, {
        type: 'line',
        data: {
            labels: chartData.map(d => d.iteration),
            datasets: [{
                label: 'Objective Function',
                data: chartData.map(d => d.OF),
                borderColor: '#2563eb',
                tension: 0.1,
                pointRadius: 0,
                borderWidth: 2
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            interaction: {
                intersect: false,
                mode: 'index'
            },
            plugins: {
                tooltip: {
                    enabled: true
                }
            },
            scales: {
                x: {
                    type: 'linear',
                    min: 0,
                    max: chartData[chartData.length-1].iteration,
                    title: {
                        display: true,
                        text: 'Iteration'
                    }
                },
                y: {
                    title: {
                        display: true,
                        text: 'Objective Function'
                    },
                    beginAtZero: true
                }
            }
        },
        plugins: [createVerticalLinePlugin()]
    });
    
    // Create Avg chart
    avgChart = new Chart(avgCanvas, {
        type: 'line',
        data: {
            labels: chartData.map(d => d.iteration),
            datasets: [{
                label: 'Average Objective Function',
                data: chartData.map(d => d.AvgOF),
                borderColor: '#2563eb',
                tension: 0.1,
                pointRadius: 0,
                borderWidth: 2
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            interaction: {
                intersect: false,
                mode: 'index'
            },
            plugins: {
                tooltip: {
                    enabled: true
                }
            },
            scales: {
                x: {
                    type: 'linear',
                    min: 0,
                    max: chartData[chartData.length-1].iteration,
                    title: {
                        display: true,
                        text: 'Iteration'
                    }
                },
                y: {
                    title: {
                        display: true,
                        text: 'Avg Objective Function'
                    },
                    beginAtZero: true
                }
            }
        },
        plugins: [createVerticalLinePlugin()]
    });
}

function createVerticalLinePlugin() {
    return {
        id: 'verticalLine',
        afterDraw: (chart) => {
            if (currentStateIndex !== undefined) {
                const meta = chart.getDatasetMeta(0);
                if (!meta.hidden) {
                    const data = chart.data.labels;
                    let closestIdx = 0;
                    let minDiff = Math.abs(data[0] - currentStateIndex);
                    
                    for (let i = 1; i < data.length; i++) {
                        const diff = Math.abs(data[i] - currentStateIndex);
                        if (diff < minDiff) {
                            minDiff = diff;
                            closestIdx = i;
                        }
                    }
                    
                    const x = meta.data[closestIdx].x;
                    const ctx = chart.ctx;
                    const yAxis = chart.scales.y;
                    
                    ctx.save();
                    ctx.beginPath();
                    ctx.moveTo(x, yAxis.top);
                    ctx.lineTo(x, yAxis.bottom);
                    ctx.lineWidth = 2;
                    ctx.strokeStyle = '#f73939';
                    ctx.stroke();
                    ctx.restore();
                }
            }
        }
    };
}