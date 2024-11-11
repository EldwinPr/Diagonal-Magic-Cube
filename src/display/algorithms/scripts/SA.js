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
        document.getElementById('InitT').textContent = data.states[0].temperature.toFixed(4);
        document.getElementById('FinalOF').textContent = data.finalOF;
        document.getElementById('TotalI').textContent = data.states[data.states.length-1].iteration;
        document.getElementById('stuck').textContent = 1 - (data.customVar / (data.states[data.states.length-1].iteration));
        document.getElementById('Duration').textContent = formatDuration(data.duration);
        document.getElementById('Comp').textContent = (100 - ((data.finalOF / data.initialOF)*100)).toFixed(4);
        
        // Create chart data
        const chartData = data.states.map(state => ({
            iteration: state.iteration,
            OF: state.OF,
            prob: state.prob
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
    document.getElementById('temp').textContent = state.temperature.toFixed(4);
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

//Chart Creation
let myChart = null;
let probChart = null;

function createChart(chartData) {
    const canvas = document.getElementById('ofChart');
    const probCanvas = document.getElementById('probChart');
    
    if (myChart instanceof Chart) {
        myChart.destroy();
    }
    if (probChart instanceof Chart) {
        probChart.destroy();
    }

    // Reset to null after destroy
    myChart = null;
    probChart = null;

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

    // Create probability chart
    const probData = chartData.map(d => {
        return {
            x: d.iteration,
            y: d.prob !== undefined ? parseFloat(d.prob.toFixed(4)) : 0
        };
    });
    
    probChart = new Chart(probCanvas, {
        type: 'scatter',
        data: {
            datasets: [{
                label: 'e^(ΔE/T)',
                data: probData,
                backgroundColor: '#10b981',
                pointRadius: 2
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
                    enabled: true,
                    callbacks: {
                        label: function(context) {
                            return `e^(ΔE/T): ${context.parsed.y.toFixed(4)}`;
                        }
                    }
                }
            },
            scales: {
                x: {
                    type: 'linear',
                    min: 0,
                    max: chartData[chartData.length-1].iteration
                },
                y: {
                    min: 0,
                    max: 1,
                    ticks: {
                        callback: function(value) {
                            return value.toFixed(4);
                        }
                    }
                }
            }
        }
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

