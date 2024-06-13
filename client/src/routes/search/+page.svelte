<script>
	import ErrorMessage from '$lib/components/ErrorMessage.svelte';
	import container from '$lib/config/bootstrap';

	// Resolves the service through the IoCContainer.
	const service = container.resolve('SearchService');

	// The current value of the searchbar.
	var searchBarValue = '';
	// The value used for the latest search.
	var searchedValue = '';
	// Variable holding the returned data from the search.
	let searchData;
	// Variable holding the aria-invalid state for the searchbar.
	let ariaInvalid = '';
	// Current page of pagination.
	let currentPage = 0;
	// Amount of items returned per search, typesense default.
	let amountPerPage = 10;
	// The calculated amount of pages.
	let totalPages = 0;

	let errorComponent;

	// Calculates amount of pages based on the returned data.
	function calculatePages() {
		if (searchData.found === 0) {
			currentPage = 0;
			totalPages = 0;
		} else {
			currentPage = 1;
			totalPages = Math.ceil(searchData.found / amountPerPage);
		}
	}

	// Performs the initial search / search based on content in the search bar.
	async function performInitialSearch() {
		try {
			if (searchBarValue.trim() === '') {
				ariaInvalid = 'true';
				return;
			} else {
				ariaInvalid = '';
				searchedValue = searchBarValue;
				currentPage = 1;
				searchData = await service.searchData(searchBarValue, currentPage);
				calculatePages();
			}
		} catch (error) {
			errorComponent.displayMessage('No data available..');
		}
	}

	// Gets specified page of last searched value, used for getting next
	// and previous page of a search.
	async function fetchPage(page) {
		if (page < 1 || page > totalPages) return;
		try {
			currentPage = page;
			searchData = await service.searchData(searchedValue, currentPage);
		} catch (error) {
			errorComponent.displayMessage('No data available..');
		}
	}

	// Reactive value for handling next and previous buttons state.
	$: hasPreviousPage = currentPage > 1;
	$: hasNextPage = currentPage < totalPages;
</script>

<div class="container">
	<ErrorMessage bind:this={errorComponent} />
	<form class="container" id="main-wrapper" on:submit|preventDefault={performInitialSearch}>
		<input
			type="search"
			name="search"
			placeholder="Search for games by name, platform, genre or publisher.."
			aria-label="Search"
			id="search-bar"
			bind:value={searchBarValue}
			aria-invalid={ariaInvalid}
		/>
	</form>

	{#if searchData}
		<p>Showing page {currentPage} of {totalPages}</p>
		<table>
			<thead>
				<tr>
					<th scope="col">Name</th>
					<th scope="col">Platform</th>
					<th scope="col">Year</th>
					<th scope="col">Genre</th>
					<th scope="col">Publisher</th>
					<th scope="col">NaSales</th>
					<th scope="col">EuSales</th>
					<th scope="col">JpSales</th>
					<th scope="col">OtherSales</th>
					<th scope="col">GlobalSales</th>
				</tr>
			</thead>
			<tbody>
				{#each searchData.videoGamesData as game}
					<tr>
						<th scope="row">{game.name}</th>
						<td>{game.platform}</td>
						<td>{game.year}</td>
						<td>{game.genre}</td>
						<td>{game.publisher}</td>
						<td>{game.naSales}</td>
						<td>{game.euSales}</td>
						<td>{game.jpSales}</td>
						<td>{game.otherSales}</td>
						<td>{game.globalSales}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
	{#if totalPages > 1}
		<div role="group" id="pagination-buttons">
			<button
				class="outline"
				disabled={!hasPreviousPage}
				on:click={() => fetchPage(currentPage - 1)}>Previous page</button
			>
			<button class="outline" disabled={!hasNextPage} on:click={() => fetchPage(currentPage + 1)}
				>Next page</button
			>
		</div>
	{/if}
</div>

<style>
	#main-wrapper {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	#search-bar {
		width: 50%;
	}
	#pagination-buttons {
		width: 30%;
	}
</style>
