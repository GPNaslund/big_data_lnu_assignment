<script>
	import { createEventDispatcher } from 'svelte';

	// The options to be displayed in the dropdown component.
	export var options;

	// The value to be displayed as a placeholder for the dropdown.
	export var displayValue;

	// Creates a dispatcher in svelte.
	const dispatch = createEventDispatcher();

	// Method for dispatching an event when selection is made.
	function handleChange(event) {
		dispatch('change', event.target.value);
	}
</script>

<select name="select" aria-label="Select" required on:change={handleChange}>
	<option selected disabled value="">{displayValue}</option>
	{#each options as option}
		<option>{option}</option>
	{/each}
</select>
