let myChart = null;

function createChart(chartData) {
    const canvas = document.getElementById('ofChart');
    
    // Destroy existing chart if it exists
    if (myChart) {
        myChart.destroy();
    }

    // Create new chart
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
                    type: 'linear', // Force linear scale
                    min: 0,
                    max: chartData[chartData.length-1].iteration, // Use actual max iteration
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
}