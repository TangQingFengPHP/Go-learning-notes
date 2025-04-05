let chart;
const ctx = document.getElementById('benchmarkChart').getContext('2d');

function renderChart(data, metric = "ns_per_op") {
    const labels = data.map(item => item.name);
    const values = data.map(item => item[metric]);

    if (chart) chart.destroy();

    chart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: labels,
            datasets: [{
                label: metric,
                data: values,
                backgroundColor: 'rgba(54, 162, 235, 0.6)'
            }]
        },
        options: {
            scales: {
                y: {
                    beginAtZero: true
                }
            }
        }
    });
}

fetch('/data')
    .then(res => res.json())
    .then(data => {
        renderChart(data);
        document.getElementById('metric-select').addEventListener('change', (e) => {
            renderChart(data, e.target.value);
        });
    });
