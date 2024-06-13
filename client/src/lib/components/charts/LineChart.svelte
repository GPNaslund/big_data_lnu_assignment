<script>
	import Chart from 'chart.js/auto';
	import { onMount } from 'svelte';

	// The chart instance created by Chart.js.
	let chart;

	// The HTMLCanvas element that holds the chart instance.
	let chartElement;

	// Object containing the chart data to be displayed.
	export let chartData;

	// The max value of the y axis.
	export let maxValue;

	// Creates the chart instance when component is mounted.
	onMount(() => {
		chart = new Chart(chartElement, {
			type: 'line',
			data: chartData,
			options: {
				scales: {
					y: {
						beginAtZero: true,
						max: maxValue
					}
				},
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

		// Destroys the chart when component is unmounted.
		return () => {
			chart.destroy();
		};
	});

	// Reactive handling of chart + chartdata for updating the visuals.
	$: if (chart && (chartData || maxValue !== undefined)) {
		chart.data = chartData;
		chart.options.scales.y.max = maxValue;
		chart.update();
	}

	// Method exposed outside of the component for updating the chart.
	export function update(newData, newMaxValue) {
		chartData = newData;
		if (newMaxValue !== undefined) {
			maxValue = newMaxValue;
		}
	}
</script>

<p>Sales are displayed in millions</p>
<canvas bind:this={chartElement}></canvas>
