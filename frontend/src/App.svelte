<script>
	// Import the router component
	import Router from 'svelte-spa-router'

	// Import the pages
	import Home from './pages/Home.svelte'
	import Name from './pages/Name.svelte'
	import Wild from './pages/Wild.svelte'
	import NotFound from './pages/NotFound.svelte'

	// Create routes definition
	const routes = {
		// Exact path
		'/': Home,

		// Using named parameters, with last being optional
		'/hello/:first/:last?': Name,

		// Wildcard parameter
		// Included twice to match both `/wild` (and nothing after) and `/wild/*` (with anything after)
		'/wild': Wild,
		'/wild/*': Wild,

		// Catch-all, must be last
		'*': NotFound
	}
</script>

<h1>svelte-spa-router sample</h1>
<h2>Basic routing</h2>

<!-- Navigation links -->
<ul>
	<li><a href="#/">Home</a></li>
	<li><a href="#/hello/svelte">Say hi!</a></li>
	<li><a href="#/wild/card">Wildcard route</a></li>
	<li><a href="#/does/not/exist">Not found</a></li>
</ul>

<!-- Show the pages -->
<Router {routes} />
