<script>
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	import DropdownSelect from './partials/DropdownSelect.svelte';
	import DropdownCheckbox from './partials/DropdownCheckbox.svelte';

	export var formParameters;

	// Variable holding the type of aggregation choosen.
	let aggregationChoice = '';
	let allAggregations = formParameters.aggregations;

	// Variable holding the type of grouping choosen.
	let groupingChoice = '';
	let allGroupings = formParameters.groupings;

	// The regions available for selection if grouping is by region.
	let allRegions = formParameters.regions;
	let selectedRegions = Array.from(allRegions);

	// The genres available for selection if grouping is by genre.
	let allGenres = formParameters.genres;

	let selectedGenres = Array.from(allGenres);

	// Variable holding the start year for timespan slider.
	let startYear = formParameters.startYear;
	let selectedStartYear = startYear;
	// Variable holding the end year for timespan slider.
	let endYear = formParameters.endYear;
	let selectedEndYear = endYear;

	// Variable holding the selected year for specific year selection.
	let selectedYear = startYear;

	let disableSubmitBtn = false;
	let regionInvalid = false;
	let genreInvalid = false;

	// Registers changes to form element values.
	function handleSelection(type, event) {
		if (type === 'aggregation') {
			aggregationChoice = event.detail;
		}

		if (type === 'grouping') {
			groupingChoice = event.detail;
		}

		if (type === 'region') {
			selectedRegions = event.detail;
			if (selectedRegions.length === 0) {
				regionInvalid = true;
				disableSubmitBtn = true;
			} else {
				regionInvalid = false;
				disableSubmitBtn = false;
			}
		}

		if (type === 'genre') {
			selectedGenres = event.detail;
			if (selectedGenres.length === 0) {
				genreInvalid = true;
				disableSubmitBtn = true;
			} else {
				genreInvalid = false;
				disableSubmitBtn = false;
			}
		}
	}

	// Reactive variable for syncing startyear and endyear sliders.
	$: if (selectedStartYear > selectedEndYear) {
		selectedEndYear = selectedStartYear;
	}

	// Method for collecting selected options and dispatching event containing the details.
	function handleButtonClick() {
		const resultObj = {
			aggregationChoice: aggregationChoice,
			groupingChoice: groupingChoice
		};

		if (aggregationChoice == 'by year') {
			resultObj.years = {
				startYear: selectedStartYear,
				endYear: selectedEndYear
			};
		}

		if (aggregationChoice == 'all games') {
			resultObj.selectedYear = selectedYear;
		}

		if (groupingChoice == 'by region') {
			resultObj.filtering = selectedRegions;
		}

		if (groupingChoice == 'by genre') {
			resultObj.filtering = selectedGenres;
		}

		dispatch('submit', resultObj);
	}
</script>

<div id="form-wrapper">
	<DropdownSelect
		displayValue="Select aggregation"
		options={allAggregations}
		on:change={(event) => handleSelection('aggregation', event)}
	/>

	{#if aggregationChoice == 'by year'}
		<div id="year-selection">
			<label>
				From: {selectedStartYear}
				<input
					type="range"
					id="start-year"
					min={startYear}
					max={endYear}
					bind:value={selectedStartYear}
				/>
			</label>
			<label>
				To: {selectedEndYear}
				<input
					type="range"
					id="end-year"
					min={startYear}
					max={endYear}
					bind:value={selectedEndYear}
				/>
			</label>
		</div>
	{/if}

	{#if aggregationChoice == 'all games'}
		<div id="year-selection">
			<label>
				Year: {selectedYear}
				<input
					type="range"
					id="selected-year"
					min={startYear}
					max={endYear}
					bind:value={selectedYear}
				/>
			</label>
		</div>
	{/if}

	{#if aggregationChoice && aggregationChoice != 'all games'}
		<DropdownSelect
			displayValue="Select grouping"
			options={allGroupings}
			on:change={(event) => handleSelection('grouping', event)}
		/>
	{/if}

	{#if groupingChoice == 'by region' && aggregationChoice == 'by year'}
		<DropdownCheckbox
			summary="Selected regions"
			options={allRegions}
			on:change={(event) => handleSelection('region', event)}
			invalid={regionInvalid}
		/>
	{/if}

	{#if groupingChoice == 'by genre' && aggregationChoice == 'by year'}
		<DropdownCheckbox
			summary="Selected genres"
			options={allGenres}
			on:change={(event) => handleSelection('genre', event)}
			invalid={genreInvalid}
		/>
	{/if}

	{#if groupingChoice || aggregationChoice == 'all games'}
		<button id="load-btn" on:click={handleButtonClick} disabled={disableSubmitBtn}>Load data</button
		>
	{/if}
</div>

<style>
	#form-wrapper {
		display: flex;
		flex-direction: column;
		gap: 5%;
		width: 20%;
	}
</style>
