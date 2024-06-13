<script>
	import { parameters } from '$lib/store';
	import { onMount } from 'svelte';
	import DoughnutChart from '$lib/components/charts/DoughnutChart.svelte';
	import LineChart from '$lib/components/charts/LineChart.svelte';
	import DashboardForm from '$lib/components/form/DashboardForm.svelte';
	import container from '$lib/config/bootstrap';
	import DoughnutChartData from '$lib/model/DoughnutChartData';
	import LineChartData from '$lib/model/LineChartData';
	import BubbleChart from '$lib/components/charts/BubbleChart.svelte';
	import BubbleChartData from '$lib/model/BubbleChartData';
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';

	// Resolves the service instance from IoCContainer.
	const dashboardService = container.resolve('DashboardService');

	// Line chart controls
	let displayLineChart = false;
	let lineChartData;
	let maxValue;
	let updateLineChart;

	// Doughnut chart controls
	let displayDoughnutChart = false;
	let doughnutChartData;
	let updateDoughnutChart;

	// Bubble chart controls
	let displayBubbleChart = false;
	let bubbleChartData;
	let updateBubbleChart;

	// The flash error component.
	let errorComponent;

	let formParameters;
	$: formParameters = $parameters;

	// Checks if parameters is stored, if not fetches them from backend.
	onMount(async () => {
		if ($parameters === null) {
			try {
				const fetchedParams = await dashboardService.getValidParameters();
				parameters.set(fetchedParams);
			} catch (error) {
				console.error(error);
				errorComponent.displayMessage('Failed to get parameters..');
			}
		}
	});
	// Queries for data and loads the appropriate chart.
	async function handleFormSubmit(event) {
		try {
			const data = await dashboardService.getChartData(event.detail);

			if (data instanceof DoughnutChartData) {
				handleDoughnutChart(data);
			}

			if (data instanceof LineChartData) {
				handleLineChart(data);
			}

			if (data instanceof BubbleChartData) {
				handleBubbleChart(data);
			}
		} catch (error) {
			errorComponent.displayMessage('No data available..');
			displayDoughnutChart = false;
			displayLineChart = false;
			displayBubbleChart = false;
		}
	}

	function handleDoughnutChart(data) {
		doughnutChartData = {
			labels: data.labels,
			datasets: data.datasets
		};
		if (displayDoughnutChart) {
			updateDoughnutChart(doughnutChartData);
		} else {
			displayLineChart = false;
			displayBubbleChart = false;
			displayDoughnutChart = true;
		}
	}

	function handleLineChart(data) {
		lineChartData = {
			labels: data.labels,
			datasets: data.datasets
		};
		maxValue = data.maxSales * 1.1;
		if (displayLineChart) {
			updateLineChart(lineChartData, maxValue);
		} else {
			displayDoughnutChart = false;
			displayBubbleChart = false;
			displayLineChart = true;
		}
	}

	function handleBubbleChart(data) {
		bubbleChartData = {
			labels: data.labels,
			datasets: data.datasets
		};
		console.log(bubbleChartData);
		if (displayBubbleChart) {
			updateBubbleChart(bubbleChartData);
		} else {
			displayDoughnutChart = false;
			displayLineChart = false;
			displayBubbleChart = true;
		}
	}
</script>

<div class="container" id="main-wrapper">
	<ErrorMessage bind:this={errorComponent} />
	{#if formParameters}
		<DashboardForm on:submit={handleFormSubmit} {formParameters} />
	{/if}
	{#if !formParameters}
		<span aria-busy="true"></span>
	{/if}
	<div id="chart-wrapper">
		{#if displayDoughnutChart}
			<DoughnutChart chartData={doughnutChartData} bind:update={updateDoughnutChart} />
		{/if}
		{#if displayLineChart}
			<LineChart chartData={lineChartData} {maxValue} bind:update={updateLineChart} />
		{/if}
		{#if displayBubbleChart}
			<div id="bubble-chart">
				<BubbleChart chartData={bubbleChartData} bind:update={updateBubbleChart} />
			</div>
		{/if}
	</div>
</div>

<style>
	#main-wrapper {
		display: flex;
		gap: 10%;
		height: 80vh;
	}
	#chart-wrapper {
		width: 50%;
	}
	#bubble-chart {
		height: 110vh;
		width: 60vw;
	}
</style>
