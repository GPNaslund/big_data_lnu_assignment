<script>
	import { createEventDispatcher } from 'svelte';
	// Creates event dispatcher inside of svelte.
	const dispatch = createEventDispatcher();

	// The text to display as a placeholder for the dropdown.
	export var summary;

	// Array of string that is displayed as options in the dropdown.
	export var options;

	// Variable for handling invalid state
	export var invalid;

	// Array holding all the selected options, used to keep track so that selected values
	// can be queried.
	let selectedOptions = options;

	// Reactive variable for setting the selected options to all the provided options when added.
	$: selectedOptions = options;

	// Method for adding/removing an option from the selected options.
	function toggleOption(option) {
		if (selectedOptions.includes(option)) {
			selectedOptions = selectedOptions.filter((o) => o !== option);
		} else {
			selectedOptions = [...selectedOptions, option];
		}
		// Dispatches event that can be listen to from outside the component.
		dispatch('change', selectedOptions);
	}
</script>

<details class="dropdown">
	<summary>{summary}</summary>
	<ul>
		{#each options as option}
			<li>
				<label>
					<input
						type="checkbox"
						name={option}
						checked={selectedOptions.includes(option)}
						on:change={() => toggleOption(option)}
						aria-invalid={invalid}
					/>
					{option}
				</label>
			</li>
		{/each}
	</ul>
</details>
