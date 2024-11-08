let myChart = null;
let probChart = null;

function createChart(chartData) {
    const canvas = document.getElementById('ofChart');
    const probCanvas = document.getElementById('probChart');
    
    // Destroy existing charts if they exist
    if (myChart) {
        myChart.destroy();
    }
    if (probChart) {
        probChart.destroy();
    }

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
        plugins: [{
            id: 'verticalLine',
            afterDraw: (chart) => {
                if (currentStateIndex !== undefined) {
                    const meta = chart.getDatasetMeta(0);
                    if (!meta.hidden) {
                        // Find the closest data point
                        const data = chart.data.labels;
                        let closestIdx = 0;
                        let minDiff = Math.abs(data[0] - currentStateIndex);
                        
                        for(let i = 1; i < data.length; i++) {
                            const diff = Math.abs(data[i] - currentStateIndex);
                            if(diff < minDiff) {
                                minDiff = diff;
                                closestIdx = i;
                            }
                        }
                        
                        if(meta.data[closestIdx]) {
                            const ctx = chart.ctx;
                            const x = meta.data[closestIdx].x;
                            const yAxis = chart.scales.y;

                            ctx.save();
                            ctx.beginPath();
                            ctx.moveTo(x, yAxis.top);
                            ctx.lineTo(x, yAxis.bottom);
                            ctx.lineWidth = 2;
                            ctx.strokeStyle = '#dc2626';
                            ctx.stroke();
                            ctx.restore();
                        }
                    }
                }
            }
        }]
    });
    
    // Then modify how we create probData
    const probData = chartData.map(d => {
        console.log("Iteration:", d.iteration, "Prob:", d.prob, "Type:", typeof d.prob);
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