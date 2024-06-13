<script>
	import { base } from '$app/paths';
</script>

<nav class="container">
	<ul>
		<li><a href={base + '/'}><strong>Video games sales stats</strong></a></li>
	</ul>
	<ul>
		<li><a href={base + '/'}>Dashboard</a></li>
		<li><a href={base + '/search'}>Search</a></li>
	</ul>
</nav>

<slot />
