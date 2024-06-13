<script>
	import { onMount } from 'svelte';
	import Chart from 'chart.js/auto';

	// The chart instance created with Chart.js.
	let chart;

	// The HTMLCanvas element that holds the chart instance.
	var chartContainer;

	// Object containing the data to be displayed in the chart.
	export var chartData;

	// Creates chart when component is mounted.
	onMount(() => {
		chart = new Chart(chartContainer, {
			type: 'doughnut',
			data: chartData,
			options: {
				responsive: true,
				maintainAspectRatio: false,
				plugins: {
					legend: {
						display: true,
						position: 'top'
					}
				}
			}
		});

		// Destroys the chart when the component unmounts.
		return () => {
			chart.destroy();
		};
	});

	// Reactive variable for updating the chart dynamically without rebuilding.
	$: if (chart && chartData) {
		chart.data = chartData;
		chart.update();
	}

	// Method for updating the chart, exposed outside of the component.
	export function update(newData) {
		chartData = newData;
	}
</script>

<p>Sales are displayed in millions</p>
<canvas bind:this={chartContainer}></canvas>
