<script>
	import Chart from 'chart.js/auto';
	import { onMount } from 'svelte';

	// The chart instance created with Chart.js.
	let chart;

	// The HTMLCanvas element holding the chart.
	let chartElement;

	// Object containing chart data to be displayed in the chart.
	export let chartData;

	// The maximum value for the y-axis.
	export let maxValue;

	// Creates chart when component is mounted.
	onMount(() => {
		chart = new Chart(chartElement, {
			type: 'bubble',
			data: chartData,
			options: {
				scales: {
					x: {
						type: 'linear',
						position: 'bottom',
						display: false
					},
					y: {
						type: 'logarithmic',
						min: 0.009,
						max: maxValue,
						title: {
							display: true,
							text: 'Total Sales'
						},
						ticks: {
							callback: function (value, index, values) {
								return Number(value.toString());
							}
						}
					}
				},
				plugins: {
					tooltip: {
						callbacks: {
							// Creates the tooltip container with info specified.
							label: function (context) {
								const data = context.raw.gameInfo;
								return [
									`Name: ${data.Name}`,
									`Year: ${data.Year}`,
									`Platform: ${data.Platform}`,
									`NA Sales: ${data.NaSales.toFixed(2)}M`,
									`EU Sales: ${data.EuSales.toFixed(2)}M`,
									`JP Sales: ${data.JpSales.toFixed(2)}M`,
									`Other Sales: ${data.OtherSales.toFixed(2)}M`,
									`Global Sales: ${data.GlobalSales.toFixed(2)}M`
								];
							}
						}
					},
					legend: {
						display: true,
						position: 'top'
					}
				}
			}
		});

		// Callback to destory the chart when component is unmounted.
		return () => {
			chart.destroy();
		};
	});

	// Reactive variable to use for updating the component.
	$: if (chart && (chartData || maxValue !== undefined)) {
		chart.data = chartData;
		chart.options.scales.y.max = maxValue;
		chart.update();
	}

	// Method for updating the chart.
	export function update(newData, newMaxValue) {
		chartData = newData;
		if (newMaxValue !== undefined) {
			maxValue = newMaxValue;
		}
	}
</script>

<p>Sales are displayed in millions</p>
<canvas bind:this={chartElement}></canvas>
