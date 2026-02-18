<script lang="ts">
	import "./layout.css";
	import "../fonts/IoskeleyMono/font.css";
	import favicon from "$lib/assets/favicon.svg";
	import { ModeWatcher } from "mode-watcher";
	import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				refetchOnWindowFocus: false,
			},
		},
	});
	let { children } = $props();
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>
<ModeWatcher />
<QueryClientProvider client={queryClient}>
	{@render children()}
</QueryClientProvider>
