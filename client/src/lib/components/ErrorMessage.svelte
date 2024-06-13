<script>
	// COMPONENT IS COPIED FROM MY OWN IMPLEMENTATION IN WT1

	import { fade } from 'svelte/transition';

	// The error message to be displayed.
	var message;
	var visible;

	export function displayMessage(newMessage) {
		message = newMessage;
		visible = true;

		setTimeout(() => {
			visible = false;
		}, 3000);
	}
</script>

{#if visible}
	<input
		readonly
		type="text"
		value={message}
		transition:fade={{ duration: 500 }}
		aria-invalid="true"
		id="error-message"
	/>
{/if}

<style>
	#error-message {
		position: absolute;
		width: auto;
		height: auto;
		left: 50%;
		transform: translate(-50%, 0);
		text-align: center;
	}
</style>
