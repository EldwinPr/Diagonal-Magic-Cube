let myChart = null;

function createChart(chartData) {
    const canvas = document.getElementById('ofChart');
    
    // Safe destroy with null checks
    if (myChart instanceof Chart) {
        myChart.destroy();
    }

    // Reset to null after destroy
    myChart = null;

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

